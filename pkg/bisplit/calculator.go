package bisplit

import (
  "bytes"
)

//Calculator handles comparisons and operations on Bisplits
//of certain dimensions
type Calculator interface {
  IsNormal(a Bisplit) bool
  Normalize(a Bisplit) Bisplit
  Reverse(a Bisplit) Bisplit

  Equivalent(a Bisplit, b Bisplit) bool
  StrictlyCompatible(a Bisplit, b Bisplit) bool
  SignedCompatible(a Bisplit, b Bisplit) int
  Compatible(a Bisplit, b Bisplit) bool
}

//ModuliSpace implements Calculator
type ModuliSpace struct {
  W uint
  B uint
}

//IsNormal returns true if the Bisplit is already in normal form
func (sp ModuliSpace) IsNormal(a Bisplit) bool {
  return a.W&(1<<sp.W) == 0
}

//Normalize puts the Bisplit in normal form
func (sp ModuliSpace) Normalize(a Bisplit) Bisplit {
  if sp.IsNormal(a) {
    return a
  }

  return sp.Reverse(a)
}

//Reverse produces the reverse of a Biplit
func (sp ModuliSpace) Reverse(a Bisplit) Bisplit {
  return Bisplit{
    W: reverse(a.W, sp.W),
    B: reverse(a.B, sp.B),
  }
}

//Equivalent returns true if both Bisplits are equivalent
func (sp ModuliSpace) Equivalent(a Bisplit, b Bisplit) bool {
  if a == b {
    return true
  }

  if a == sp.Reverse(b) {
    return true
  }

  return false
}

//StrictlyCompatible returns true if a and b are strictly compatible
func (sp ModuliSpace) StrictlyCompatible(a Bisplit, b Bisplit) bool {
  diffMask := a.W ^ b.W

  if diffMask == 0 {
    diffMask = a.B ^ b.B

    //return whether all the bits that differ between them differ in the same way
    diffVals := b.B & diffMask
    return (diffVals == diffMask) || (diffVals == 0)
  }

  diffVals := b.W & diffMask
  if (diffVals != diffMask) && (diffVals != 0) {
    return false
  }

  diffMask = a.B ^ b.B
  diffValsB := b.B & diffMask
  if diffVals == 0 {
    return diffValsB == 0
  }

  return diffValsB == diffMask
}

//SignedCompatible returns true if a and b are signed compatible
func (sp ModuliSpace) SignedCompatible(a Bisplit, b Bisplit) int {
  if sp.StrictlyCompatible(a, b) {
    return 1
  }

  if sp.StrictlyCompatible(a, sp.Reverse(b)) {
    return -1
  }

  return 0
}

//Compatible returns true if a and b are compatible
func (sp ModuliSpace) Compatible(a Bisplit, b Bisplit) bool {
  return sp.SignedCompatible(a, b) != 0
}

//String provides a string representation of this Bisplit
func (sp ModuliSpace) String(a Bisplit) string {
  var buf bytes.Buffer
  bitsIntoBuffer(a.W, sp.W, &buf)
  bitsIntoBuffer(a.B, sp.B, &buf)
  return buf.String()
}

func bitsIntoBuffer(val, size uint, buf *bytes.Buffer) {
  buf.WriteString("{")
  if size > 0 {
    for i := uint(0); i < size; i++ {
      if val%2 == 1 {
        buf.WriteString("1")
      } else {
        buf.WriteString("0")
      }
      val >>= 1
    }
  }
  buf.WriteString("}")
}

//zero out everything except the least significant l bits
func scope(a uint, l uint) uint {
  return a & ((1 << l) - 1)
}

func reverse(a uint, l uint) uint {
  return scope(^a, l)
}
