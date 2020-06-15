package structfieldorder

type MyStruct struct {
	myBool   bool    // 1 byte
	myFloat float64  // 8 bytes
	myInt  int32     // 4 bytes
}


type MyStructOptimized struct {
	myFloat float64 // 8 bytes
	myBool  bool    // 1 byte
	myInt   int32   // 4 bytes
}
