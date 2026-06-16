class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        counter = {}
        for i in s:
            counter[i] = counter.get(i, 0) + 1 

        for x in t:
            if x not in counter:
                return False
            counter[x] -= 1

        for count in counter.values():
            if count != 0:
                return False

        return True

    
        