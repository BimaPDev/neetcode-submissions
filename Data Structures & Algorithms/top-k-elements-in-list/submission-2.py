class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        
        hashSet = {} #dictionary

        #Iterate throught the list
        for i in range(len(nums)):
            #for each number you will append to the dictionary
            if nums[i] in hashSet:
                hashSet[nums[i]] += 1
            else:
                hashSet[nums[i]] = 1
        
        sortedSet = sorted(hashSet, key=lambda x: hashSet[x], reverse= True)
        return sortedSet[:k]
                
             