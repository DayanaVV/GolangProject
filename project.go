package numberutil
 
import "strconv"
 
//Convert Decimal to Binary
func DecimalToBinary(decimal int64) string {
    binary := strconv.FormatInt(decimal, 2)
    return binary
}
