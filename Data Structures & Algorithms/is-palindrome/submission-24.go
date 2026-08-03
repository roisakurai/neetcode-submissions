func isPalindrome(s string) bool {
    lower := strings.ToLower(s)
    output := ""

    for i := range lower {
        if (lower[i] >= 'a' && lower[i] <= 'z') || (lower[i] >= '0' && lower[i] <='9') {
            output = output + string(lower[i])
        } 
    }

    fmt.Println(output)

    half := len(output)/2

    for j := range half {
        if output[j] != output[len(output)-1-j] {
            return false
        }
    }

    return true


}
