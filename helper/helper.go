package helper

var version = "1.0.0"      // tidak bisa diakses di luar package
var Application = "golang" // bisa diakses di luar package

func SayHello(name string) string {
	return "Hello " + name
}
