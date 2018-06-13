package bisplit

/*
 * ENCODING SPLITS AS INTS:
 * The ints in the Bisplit struct can only be understood as splits
 * with the context provided by a Bispliterator, which knows how
 * many bits to consider.
 *
 * TERMINOLOGY:
 * A split on n elements is a division of {1, 2, 3, ..., n} into two
 *   indistinguishable subsets, called right and left
 * A strict split is the same thing, but the subsets are distinguishable.
 *   NOTE: there are twice as many strict splits as splits.
 * A bicolored split or bisplit on W,B elements is a split on W+B with the
 *   property that both sides contain one of the white elements and one of
 *   the black.
 *
 * Notice that a split is equivalent to its own "reverse".
 * Also notice that a bicolored split is the same thing as two splits, one
 *   on W and one on B, with an equivalence relation that requires us to
 *   reverse *both* splits together.
 *
 * Since the bits of an int can represent a strict split, we use this last
 *   observation to represent bisplits as pairs of ints. From bwtween the
 *   equivalent representations of each bisplit, we choose the one whose
 *   R has a zero in the nth bit - this one always has the smaller numeric
 *   value.
 *
 * Consider a 3,3 bisplit on the elements {1, 2, 3, 4, 5, 6}, with 124 in
 *   the "off" group and 356 un the "on" group. It's normal form is:
 * Bisplit{
 *   W: 0x011
 *   B: 0x100
 * }
 * This could also be written in non-normal form:
 * Bisplit{
 *   W: 0x100
 *   B: 0x011
 * }
 *
 */

//Bisplit represents a bicolored split on len(R) right elements and len(L) left elements
type Bisplit struct {
  W uint
  B uint
}

//New initializes a new Bisplit
func New(w []bool, b []bool) Bisplit {
  return Bisplit{
    W: listToInt(w),
    B: listToInt(b),
  }
}

func listToInt(a []bool) uint {
  var n uint
  for i, val := range a {
    if val {
      n += (1 << uint(i))
    }
  }

  return n
}
