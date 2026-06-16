class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        hashmap = {}
        for i in range(len(nums)):
            needeNunber = target - nums[i]
            if needeNunber in hashmap:
                return [hashmap[needeNunber],i]
            hashmap[nums[i]] = i