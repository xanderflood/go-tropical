package bisplit

//PermutationCycleType two nonincreasing sequence of unsigned
//integers, representing the white and black cycle lengths
//of a bicolored permutation. To use along with a moduli space
//to do calculations, ms.W should equal the sum of pct.W, and
//same for the B fields.
type PermutationCycleType struct {
  W []uint
  B []uint
}
