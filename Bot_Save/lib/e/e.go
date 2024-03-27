package e
import
"fmt"

func Wrap(msg string, err error) error{
	return  fmt.Errorf(format: "%s: %w", msg , err)
}

func WrapIfErr(msg string, err error) error{
	if err==nil{
		return nill
	}
	return Wrap(msg, err)
}