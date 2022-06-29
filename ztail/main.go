package main

import (
	"fmt"
	"os"
)

func main() {
	args := []string(os.Args)
	//fmt.Println(signPriority(args))
	// args := []string{"0", "-c", "18446744073709551615"}
	if len(args) == 2 {
		if args[1] == "-c" {
			TailOp() // if there are no arguments
		} else if IsANumber(args[1]) == false { //tail <filename> starts
			PrintBytes(args, args[1])
			if args[len(args)-1] == "-" {
				infiniteInput()
			}
		} else if args[1] == "-" {
			infiniteInput2()
		}
	} else if (len(args) == 1 || NumOfWords(args) == 0 && IsANumber(args[len(args)-1]) != true) || (len(args) == 3 && args[1] == "-c" && args[2] == "+0") { //inifinite input starts here
		data := []byte{0, 0, 0, 0, 0, 0, 0, 0}
		if args[len(args)-1] == "-c" && NumOfWords(args) == 0 {
			TailOp()
		} else {
			for {
				// data := make([]byte, 8)

				n, err := os.Stdin.Read(data)
				if err == nil && n > 0 {
				} else {
					break
				}
			} //infinite input ends here
		}

	} else if len(args) == 3 && args[1] == "-c" && args[2] == "+" {
		InvBytes(args[2])
	} else if len(args) > 2 { //if there are more than 2 input arguments:
		if cCount(args) == 1 { //if there is only one -c:
			if lastC(args) != len(args)-1 {
				if IsANumber(args[lastC(args)+1]) == true {
					errCount1 := 0 
					errCount2 := 0
					errCount5 := 0
					for i := 1; i < len(args); i++ {
						errCount5 = 0
						if args[i] != "-c" {
							if IsANumber(args[i]) == false || (IsANumber(args[i])) == true && args[i-1] != "-c" {
								bytes, err := os.ReadFile(args[i])
								if err == nil {
									if errCount1 > errCount2 {
										errCount2++
									} else {
										errCount1++
									}
								} else {
									errCount1--
									errCount2++
								}
								if err == nil {
									errCount5++
								}
								if errCount1 == errCount2 || (errCount5 > 0 && args[i] != args[1] && cCount(args) > 1) {
									fmt.Printf("\n")
								}
								if args[lastC(args)+1] == "+0" {
									PrintBytes(args, args[i])
									if args[len(args)-1] == "-" {
										infiniteInput()
									}
								} else {
									if uint64(len(bytes)) < CNum(args[lastC(args)+1]) && (NumValidity(args[lastC(args)+1]) == "minus" || NumValidity(args[lastC(args)+1]) == "validminus" || NumValidity(args[lastC(args)+1]) == "no sign") {
										PrintBytes(args, args[i])
									} else if NumValidity(args[lastC(args)+1]) == "plus" {
										errCount3 := 0 // these are for the fkin newlines between the arguments with no errors
										errCount4 := 0
										_, err := os.ReadFile(args[i])
										if err != nil {
											if os.IsNotExist(err) {
												CantOpen(args[i])
											}
											errCount3--
											errCount4++
										} else {
											if errCount3 > errCount4 {
												errCount4++
											} else {
												errCount3++
											}
											if uint64(len(bytes)) < CNum(args[lastC(args)+1]) {
												if NumOfWords(args) != 1 {
													Title(args[i])
												}
											} else {
												if NumOfWords(args) != 1 {
													Title(args[i])
												}
												fmt.Printf(string(bytes[CNum(args[lastC(args)+1])-1:]))
												if args[len(args)-1] == "-" {
													infiniteInput()
												}
											}
										}
										if errCount1 == errCount2 || (errCount5 > 0 && args[i] != args[1] && cCount(args) > 1 && errCount1 != errCount5) {
											fmt.Printf("\n")
										}
									} else if NumValidity(args[lastC(args)+1]) == "minus" || NumValidity(args[lastC(args)+1]) == "validminus" || NumValidity(args[lastC(args)+1]) == "no sign" {
										if NumOfWords(args) != 1 {
											Title(args[i])
										}
										fmt.Printf(string(bytes[uint64(len(bytes))-CNum(args[lastC(args)+1]):]))
										if args[len(args)-1] == "-" {
											infiniteInput()
										}
									} else if NumValidity(args[lastC(args)+1]) == "invalid" {
										InvBytes(args[lastC(args)+1])
									} else {
									}
								}
							}
						}
					}
				} else if IsANumber(args[lastC(args)+1]) == false { // invalid number of bytes when there is a word past -c and not a number
					if lastC(args) == len(args)-1 {
						TailOp()
					} else {
						InvBytes(args[lastC(args)+1])
					}
				}
			} else { // if the last argument is -c and there is no number past it
				TailOp()
			}
		} else if cCount(args) > 1 {
			if lastC(args) != len(args)-1 {
				cinvalid, index := cCheck(args)
				if cinvalid != 0 && index != 0 {
					InvBytes(args[index+1])
				} else if cinvalid == 0 && index == 0 {
					errCount1 := 0 // these are for the fkin newlines between the arguments with no errors
					errCount2 := 0
					errCount5 := 0
					for i := 1; i < len(args); i++ {
						errCount5 = 0
						if args[i] != "-c" && IsANumber(args[i]) == false || (IsANumber(args[i]) == true && args[i-1] != "-c") {
							bytes, err := os.ReadFile(args[i])
							if err == nil {
								if errCount1 > errCount2 {
									errCount2++
								} else {
									errCount1++
								}
							} else {
								errCount1--
								errCount2++
							}
							if err == nil {
								errCount5++
							}
							if errCount1 == errCount2 || (errCount5 > 0 && args[i] != args[1] && cCount(args) > 1 && errCount1 != errCount5) {
								fmt.Printf("\n")
							}
							if args[lastC(args)+1] == "+0" {
								PrintBytes(args, args[i])
								if args[len(args)-1] == "-" {
									infiniteInput()
								}
							} else {
								if (NumValidity(args[lastC(args)+1]) == "minus" || NumValidity(args[lastC(args)+1]) == "validminus" || NumValidity(args[lastC(args)+1]) == "no sign") && signPriority(args) != "plus" {
									if uint64(len(bytes)) < CNum(args[lastC(args)+1]) {
										PrintBytes(args, args[i])
										if args[len(args)-1] == "-" {
											infiniteInput()
										}
									} else {
										if NumOfWords(args) != 1 {
											Title(args[i])
										}
										fmt.Printf(string(bytes[uint64(len(bytes))-CNum(args[lastC(args)+1]):]))
										if args[len(args)-1] == "-" {
											infiniteInput()
										}
									}
								} else if NumValidity(args[lastC(args)+1]) == "invalid" {
									InvBytes(args[lastC(args)+1])
								} else if NumValidity(args[lastC(args)+1]) == "plus" || signPriority(args) == "plus" {
									errCount3 := 0 // these are for the fkin newlines between the arguments with no errors
									errCount4 := 0
									_, err := os.ReadFile(args[i])
									if err != nil {
										if os.IsNotExist(err) {
											CantOpen(args[i])
										}
										errCount3--
										errCount4++
									} else {
										if errCount3 > errCount4 {
											errCount4++
										} else {
											errCount3++
										}
										if uint64(len(bytes)) < CNum(args[lastC(args)+1]) {
											if NumOfWords(args) != 1 {
												Title(args[i])
											}
										} else {
											if NumOfWords(args) != 1 {
												Title(args[i])
											}
											if signPriority(args) == "plus" && args[lastC(args)+1] == "0" {
												PrintBytes(args, args[i])
												if args[len(args)-1] == "-" {
													infiniteInput()
												}
											} else {
												fmt.Printf(string(bytes[CNum(args[lastC(args)+1])-1:]))
												if args[len(args)-1] == "-" {
													infiniteInput()
												}
											}

										}
									}
									if errCount3 == errCount4 {
										fmt.Printf("\n")
									}

								}
							}
						} else if args[i] == "-c" && args[i] != args[len(args)-1] && IsANumber(args[len(args)-1]) == true {
							i = i + 1
							continue
						}
					}
				}
			} else {
				if cCount(args) == 1 || args[len(args)-1] == "-c" {
					TailOp()
				} else {
					InvBytes("c")
				}
			}
		}

	}
	if NumOfWords(args) == 0 && cCount(args) >= 1 && IsANumber(args[len(args)-1]) {
		for {
			data := make([]byte, 8)

			n, err := os.Stdin.Read(data)
			if err == nil && n > 0 {
			} else {
				break
			}
		} //infinite input ends here
	}
}

func infiniteInput() {
	data := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	os.Stdout.WriteString("\n")
	os.Stdout.WriteString("==> ")
	os.Stdout.WriteString("standard input")
	os.Stdout.WriteString(" <==")
	os.Stdout.WriteString("\n")
	for {
		// data := make([]byte, 8)

		n, err := os.Stdin.Read(data)
		if err == nil && n > 0 {
		} else {
			break
		}
	} //infinite input ends here
}

func infiniteInput2() {
	data := []byte{0, 0, 0, 0, 0, 0, 0, 0}
	for {
		// data := make([]byte, 8)

		n, err := os.Stdin.Read(data)
		if err == nil && n > 0 {
		} else {
			break
		}
	} //infinite input ends here
}

func signPriority(args []string) string {
	for i := 1; i < len(args)-1; i++ {
		if cCount(args) > 1 {
			if NumValidity(args[i+1]) == "plus" {
				return "plus"
			}
		}
	}
	return "minus"
}

func cCheck(args []string) (int, int) {
	cinvalid := 0
	index := 0
	for i := 0; i < len(args)-1; i++ {
		if args[i] == "-c" {
			if IsANumber(args[i+1]) != true {
				cinvalid++
				if index == 0 {
					index = i
				}
			}
		}
	}
	if args[len(args)-1] == "-c" {
		cinvalid++
	}
	return cinvalid, index
}

func NumOfWords(args []string) int {
	NumOfWords := 0
	if cCount(args) == 1 {
		for i := 1; i < len(args); i++ {
			if args[i] != "-c" {
				if args[i-1] != "-c" {
					NumOfWords++
				}
			}
		}
	} else if cCount(args) > 1 {
		for i := 1; i < len(args); i++ {
			if args[len(args)-1] != "-c" {
				if IsANumber(args[lastC(args)+1]) == true {
					if args[i] != "-c" {
						if IsANumber(args[i]) == false || args[i-1] != "-c" {
							NumOfWords++
						}
					}
				}
			}

		}
	}
	return NumOfWords
}

func ValidWords(args []string) []string {
	var ValidWords []string
	if cCount(args) == 1 {
		for i := 1; i < len(args); i++ {
			if args[i] != "-c" {
				if args[i-1] != "-c" {
					ValidWords = append(ValidWords, args[i])
				}
			}
		}
	} else if cCount(args) > 1 {
		for i := 1; i < len(args); i++ {
			if IsANumber(args[lastC(args)+1]) == true {
				if args[i] != "-c" {
					if IsANumber(args[i]) == false || args[i-1] != "-c" {
						ValidWords = append(ValidWords, args[i])
					}
				}
			}
		}
	}
	return ValidWords
}

func NumValidity(arg string) string {
	var signArr []byte
	for i := 0; i < len(arg); i++ {
		if arg[i] == 43 || arg[i] == 45 {
			signArr = append(signArr, arg[i])
		}
	}
	if len(signArr) == 2 && signArr[len(signArr)-1] == '+' && signArr[len(signArr)-2] == '-' {
		return "validminus"
	} else if len(signArr) == 2 && signArr[len(signArr)-1] == '-' && signArr[len(signArr)-2] == '+' {
		return "invalid"
	} else if len(signArr) == 2 && signArr[len(signArr)-1] == '-' && signArr[len(signArr)-2] == '-' {
		return "invalid"
	} else if len(signArr) > 2 {
		return "invalid"
	} else if len(signArr) == 1 && signArr[0] == 45 {
		return "minus"
	} else if len(signArr) == 1 && signArr[0] == 43 {
		return "plus"
	} else if len(signArr) == 0 {
		return "no sign"
	}
	return "no sign"
}

func PrintBytes(args []string, arg string) {
	bytes, err := os.ReadFile(arg)
	if err != nil {
		if os.IsNotExist(err) {
			CantOpen(arg)
		}
	} else {
		if NumOfWords(args) != 1 {
			Title(arg)
		}
		fmt.Printf(string(bytes))
	}
}

func cCount(args []string) int { //counts -c number
	n := 0
	for i := 1; i < len(args); i++ {
		if args[i] == "-c" {
			n++
		}
	}
	return n
}

func NumCount(args []string) int { //counts the number of numbers
	n := 0
	for i := 1; i < len(args); i++ {
		if IsANumber(args[i]) == true {
			n++
		}
	}
	return n
}

func lastC(args []string) int { //returns the index of the last -c
	lastC := 0
	for i := 1; i < len(args)-1; i++ {

		if args[i] == "-c" {
			lastC = i
		}
	}
	if args[len(args)-1] == "-c" {
		lastC = len(args) - 1
	}
	return lastC
}

func Title(arg string) {
	os.Stdout.WriteString("==> ")
	os.Stdout.WriteString(arg)
	os.Stdout.WriteString(" <==")
	os.Stdout.WriteString("\n")
}

func IsANumber(n string) bool {
	number := 0
	for _, letter := range n {
		if ((letter) > 57 || (letter) < 48) && letter != 45 && letter != 43 {
			number++
		}
	}

	if number == 0 {
		return true
	}
	return false
}

func RecursivePowerZtail(nb uint64, power uint64) uint64 {
	if power == 0 {
		return 1
	} else if power < 0 {
		return 0
	}
	return nb * RecursivePowerZtail(nb, power-1)
}

func InvBytes(arg string) {
	os.Stdout.WriteString("tail: invalid number of bytes: '")
	os.Stdout.WriteString(arg)
	os.Stdout.WriteString("'\n")
	os.Exit(0)
}

//18446744073709551615 max uint64
func CNum(arg string) uint64 {
	c := []rune(arg)
	var result uint64
	result = 0
	var resultOld uint64
	resultOld = 0
	counter := 0
	if c[0] == 45 || c[0] == 43 {
		for j := 1; j < len(c); j++ {
			if c[j] != 43 && c[j] != 45 {
				if uint64(c[j])-48 == 0 {
					counter++
				}
				//fmt.Println(RecursivePowerZtail(10, (uint64(len(c)-j))-uint64(1)))
				result = resultOld + (uint64(c[j])-48)*RecursivePowerZtail(10, (uint64(len(c)-j))-uint64(1))
				if result < resultOld || (counter > 19 && uint64(c[1])-48 == 1) || (counter > 18 && uint64(c[1])-48 > 1) {
					os.Stdout.WriteString("tail: invalid number of bytes: '")
					os.Stdout.WriteString(arg)
					os.Stdout.WriteString("': Value too large for defined data type\n")
					os.Exit(0)
				}
				resultOld = result
			}
		}
	} else {
		for j := 0; j < len(c); j++ {
			if c[j] != 43 && c[j] != 45 {
				if uint64(c[j])-48 == 0 {
					counter++
				}
				result = resultOld + (uint64(c[j])-48)*RecursivePowerZtail(10, (uint64(len(c)-j))-uint64(1))
				if result < resultOld || (counter > 19 && uint64(c[0])-48 == 1) || (counter > 18 && uint64(c[0])-48 > 1) {
					os.Stdout.WriteString("tail: invalid number of bytes: '")
					os.Stdout.WriteString(arg)
					os.Stdout.WriteString("': Value too large for defined data type\n")
					os.Exit(0)
				}
				resultOld = result
			}

		}
	}
	return result
}

func TailOp() {
	os.Stdout.WriteString("tail: option requires an argument -- 'c'\nTry 'tail --help' for more information.\n")
}

func TailOpInv(args []string) {
	os.Stdout.WriteString("tail: option used in invalid context -- ")
	os.Stdout.WriteString(args[len(args)-1][1:])
	os.Stdout.WriteString("\n")
}

func CantOpen(arg string) {
	os.Stdout.WriteString("tail: cannot open '")
	os.Stdout.WriteString(arg)
	os.Stdout.WriteString("' for reading: No such file or directory\n")
}

func TooLarge(arg string) {
	os.Stdout.WriteString("tail: invalid number of bytes: '")
	os.Stdout.WriteString(arg)
	os.Stdout.WriteString("': Value too large for defined data type\n")
}