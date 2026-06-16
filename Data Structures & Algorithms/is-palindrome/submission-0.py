import re
class Solution:
    def isPalindrome(self, s: str) -> bool:
        t = re.sub(r'\W+', '', s)
        b = t.lower()
        i, j = 0, len(b) - 1;
        while i < j:
            if b[i] != b[j]:
                return False
            else:
                i += 1
                j -= 1
        return True
