func lengthOfLastWord(s string) int {
    output := 0

    if len(s) == 1 {
        return 1
    }

    for i := len(s)-1; i >= 0; i--{
        if string(s[i]) != " " {
            output++
        
        if i != 0 {
            if string(s[i-1]) == " " {
            return output
            }
            }

        }
    }

    return output

}