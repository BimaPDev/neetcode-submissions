from typing import List

class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        if not nums:
            return []

        freq = {}
        for x in nums:
            freq[x] = freq.get(x, 0) + 1

        maxFreq = max(freq.values())
        buckets = [[] for _ in range(maxFreq + 1)]
        for num, count in freq.items():
            buckets[count].append(num)

        res = []
        for f in range(maxFreq, 0, -1):
            for num in buckets[f]:
                res.append(num)
                if len(res) == k:
                    return res
        return res
