package myPkg

type Configuration struct {
	Nama string 
	Usia string 
}

var Data = Configuration{Nama:"Lintang", Usia:"23"}

func Call() string {
	return "hello!"
}