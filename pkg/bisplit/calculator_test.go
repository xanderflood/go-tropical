package bisplit_test

import (
  "github.com/xanderflood/go-tropical/pkg/bisplit"

  . "github.com/onsi/ginkgo"
  . "github.com/onsi/gomega"
)

var _ = Describe("Calculator", func() {
  var (
    W uint
    B uint

    sp bisplit.Calculator
  )

  Describe("Reverse", func() {
    var (
      a   bisplit.Bisplit
      ret bisplit.Bisplit
    )

    JustBeforeEach(func() {
      sp = bisplit.ModuliSpace{W: W, B: B}
      ret = sp.Reverse(a)
    })

    Context("in the first example", func() {
      BeforeEach(func() {
        W = 3
        B = 4
        a = bisplit.New(
          []bool{true, false, false},
          []bool{false, true, false, false},
        )
      })

      It("should correctly reverse the bisplit", func() {
        expected := bisplit.New(
          []bool{false, true, true},
          []bool{true, false, true, true},
        )

        Expect(ret).To(Equal(expected))
      })
    })
  })

  Describe("StrictlyCompatible", func() {
    var (
      a   bisplit.Bisplit
      b   bisplit.Bisplit
      ret bool
    )

    JustBeforeEach(func() {
      sp = bisplit.ModuliSpace{W: W, B: B}
      ret = sp.StrictlyCompatible(a, b)
    })

    Context("in the first example", func() {
      BeforeEach(func() {
        W = 3
        B = 4
        a = bisplit.New(
          []bool{true, false, false},
          []bool{false, true, false, false},
        )
        b = bisplit.New(
          []bool{true, true, true},
          []bool{false, true, false, false},
        )
      })

      It("should be strictly compatible", func() {
        Expect(ret).To(BeTrue())
      })
    })

    Context("in the second example", func() {
      BeforeEach(func() {
        W = 3
        B = 4
        a = bisplit.New(
          []bool{true, false, false},
          []bool{false, true, false, false},
        )
        b = bisplit.New(
          []bool{true, true, true},
          []bool{false, false, false, false},
        )
      })

      It("should not be strictly compatible", func() {
        Expect(ret).To(BeFalse())
      })
    })

    Context("in the third example", func() {
      BeforeEach(func() {
        W = 3
        B = 4
        a = bisplit.New(
          []bool{true, false, false},
          []bool{false, true, false, false},
        )
        b = bisplit.New(
          []bool{true, true, true},
          []bool{true, true, true, true},
        )
      })

      It("should be strictly compatible", func() {
        Expect(ret).To(BeTrue())
      })
    })

    Context("in the fourth example", func() {
      BeforeEach(func() {
        W = 3
        B = 4
        a = bisplit.New(
          []bool{true, false, false},
          []bool{false, true, false, false},
        )
        b = bisplit.New(
          []bool{false, false, false},
          []bool{false, true, false, true},
        )
      })

      It("should not be strictly compatible", func() {
        Expect(ret).To(BeFalse())
      })
    })
  })

  Describe("SignedCompatible", func() {
    var (
      a   bisplit.Bisplit
      b   bisplit.Bisplit
      ret int
    )

    JustBeforeEach(func() {
      sp = bisplit.ModuliSpace{W: W, B: B}
      ret = sp.SignedCompatible(a, b)
    })

    Context("in the first example", func() {
      BeforeEach(func() {
        W = 3
        B = 4
        a = bisplit.New(
          []bool{true, false, false},
          []bool{false, true, false, false},
        )
        b = bisplit.New(
          []bool{true, true, true},
          []bool{false, true, false, false},
        )
      })

      It("should be positively compatible", func() {
        Expect(ret).To(Equal(1))
      })
    })

    Context("in the second example", func() {
      BeforeEach(func() {
        W = 3
        B = 4
        a = bisplit.New(
          []bool{true, false, false},
          []bool{false, true, false, false},
        )
        b = bisplit.New(
          []bool{true, true, true},
          []bool{true, false, true, true},
        )
      })

      It("should be negatively compatible", func() {
        Expect(ret).To(Equal(-1))
      })
    })

    Context("in the third example", func() {
      BeforeEach(func() {
        W = 3
        B = 4
        a = bisplit.New(
          []bool{true, false, false},
          []bool{false, true, false, false},
        )
        b = bisplit.New(
          []bool{false, false, true},
          []bool{true, false, false, true},
        )
      })

      It("should be negatively compatible", func() {
        Expect(ret).To(Equal(-1))
      })
    })

    Context("in the fourth example", func() {
      BeforeEach(func() {
        W = 3
        B = 4
        a = bisplit.New(
          []bool{true, false, false},
          []bool{false, true, false, false},
        )
        b = bisplit.New(
          []bool{true, true, true},
          []bool{false, false, true, true},
        )
      })

      It("should not be compatible", func() {
        Expect(ret).To(Equal(0))
      })
    })
  })

  Describe("Compatible", func() {
    var (
      a   bisplit.Bisplit
      b   bisplit.Bisplit
      ret bool
    )

    JustBeforeEach(func() {
      sp = bisplit.ModuliSpace{W: W, B: B}
      ret = sp.Compatible(a, b)
    })

    Context("in the first example", func() {
      BeforeEach(func() {
        W = 3
        B = 4
        a = bisplit.New(
          []bool{true, false, false},
          []bool{false, true, false, false},
        )
        b = bisplit.New(
          []bool{true, true, true},
          []bool{false, true, false, false},
        )
      })

      It("should be compatible", func() {
        Expect(ret).To(BeTrue())
      })
    })

    Context("in the second example", func() {
      BeforeEach(func() {
        W = 3
        B = 4
        a = bisplit.New(
          []bool{true, false, false},
          []bool{false, true, false, false},
        )
        b = bisplit.New(
          []bool{true, true, true},
          []bool{true, false, true, true},
        )
      })

      It("should be compatible", func() {
        Expect(ret).To(BeTrue())
      })
    })

    Context("in the third example", func() {
      BeforeEach(func() {
        W = 3
        B = 4
        a = bisplit.New(
          []bool{true, false, false},
          []bool{false, true, false, false},
        )
        b = bisplit.New(
          []bool{false, false, true},
          []bool{true, false, false, true},
        )
      })

      It("should be compatible", func() {
        Expect(ret).To(BeTrue())
      })
    })

    Context("in the fourth example", func() {
      BeforeEach(func() {
        W = 3
        B = 4
        a = bisplit.New(
          []bool{true, false, false},
          []bool{false, true, false, false},
        )
        b = bisplit.New(
          []bool{true, true, true},
          []bool{false, false, true, true},
        )
      })

      It("should not be compatible", func() {
        Expect(ret).To(BeFalse())
      })
    })
  })

  Describe("ApplyCanonical", func() {
    var (
      a   bisplit.Bisplit
      pct bisplit.PermutationCycleType
      ret bisplit.Bisplit
    )

    JustBeforeEach(func() {
      sp = bisplit.ModuliSpace{W: W, B: B}
      ret = sp.ApplyCanonical(a, pct)
    })

    Context("when W=3", func() {
      BeforeEach(func() { W = 3 })

      Context("when B=5", func() {
        BeforeEach(func() { B = 5 })

        Context("with the first permutation", func() {
          BeforeEach(func() {
            pct = bisplit.PermutationCycleType{
              W: []uint{1, 2},
              B: []uint{2, 1, 2},
            }
          })

          Context("with the first bisplit", func() {
            BeforeEach(func() {
              a = bisplit.Bisplit{
                W: 5,
                B: 21,
              }
            })

            It("permutes correctly", func() {
              Expect(ret).To(Equal(bisplit.Bisplit{
                W: 3,
                B: 14,
              }))
            })
          })

          Context("with the second bisplit", func() {
            BeforeEach(func() {
              a = bisplit.New(
                []bool{false, true, false},
                []bool{true, true, false, false, true},
              )
            })

            It("permutes correctly", func() {
              Expect(ret).To(Equal(bisplit.New(
                []bool{false, false, true},
                []bool{true, true, false, true, false},
              )))
            })
          })
        })
      })
    })
  })
})
