func mergeAlternately(word1 string, word2 string) string {
    s := strings.Builder{}
    i := 0
    
    // Allocate enough space to merged string
    s.Grow(len(word1) + len(word2))
    
    // Loop through word1 and word 2
    for i < len(word1) || i < len(word2) {
        
        // Add letter from word1 if it exists
        if i < len(word1) {
            s.WriteByte(word1[i])
        }
    
        // Add letter from word2 if it exists
        if i < len(word2) {
            s.WriteByte(word2[i])
        }
        
        i++
    }
    
    // Return merged string
    return s.String()
}