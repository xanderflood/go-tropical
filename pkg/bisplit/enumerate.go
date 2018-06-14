package bisplit

//Enumerator a channel for enumerating Bisplits
type Enumerator <-chan Bisplit

//EnumeratorFactory a function for producing enumerators
type EnumeratorFactory func(W, B int) Enumerator

//Enumerate the correct sequentiql enumerator
func Enumerate(W, B uint) Enumerator {
  bisplitChan := make(chan Bisplit, 3)

  one := uint(1)

  go func() {
    //If either one is zero, send nothing and close the channel
    if W > 1 && B > 1 {
      //only shift W-1, since we want the largest possible _normalized_ value
      wMax := (one << (W - 1)) - 1
      bMax := (one << B) - 2

      for w := one; w <= wMax; w++ {
        for b := one; b <= bMax; b++ {
          bisplitChan <- Bisplit{W: w, B: b}
        }
      }
    }

    close(bisplitChan)
  }()

  return bisplitChan
}
