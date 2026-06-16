from typing import List

class Solution:
    def encode(self, strs: List[str]) -> str:
        return ",".join(f"{len(s)}#{s}" for s in strs)

    def decode(self, s: str) -> List[str]:
        res, i, n = [], 0, len(s)
        while i < n:
            if s[i] == ',':  # skip separator inserted by encode
                i += 1
                continue
            j = i
            while j < n and s[j] != '#':  # bound check
                j += 1
            length = int(s[i:j])
            start = j + 1
            end = start + length
            res.append(s[start:end])
            i = end
        return res
