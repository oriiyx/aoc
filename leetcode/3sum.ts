function threeSum(nums: number[]): number[][] {
    nums.sort((a, b) => a - b);

    const results: number[][] = [];

    for (let i = 0; i < nums.length; i++) {
        if (i > 0 && nums[i] === nums[i - 1]) {
            continue; // skip duplicated elements
        }

        let left = i + 1;
        let right = nums.length - 1;

        while (left < right) {
            const sum = nums[i] + nums[left] + nums[right];

            if (sum === 0) {
                results.push([nums[i], nums[left], nums[right]]);
                
                // optimization part - we want to remove every number in the sorted array that has the same value as we just pushed in the array - this way we just skip every other duplicated integer in the array
                while (left < right && nums[left] === nums[left + 1]) {
                    left++; // skip duplicates                    
                }

                while (left < right && nums[right] === nums[right - 1]) {
                    right--; // skip duplicates
                }

                left++;
                right--;

            } else if (sum < 0) {
                left++; // because we sorted the array we know that we need a bigger integer becazse sum is smaller than 0
            } else {
                right--; // because we sorted the array and we know that our sum is larger than 0 we want to get a smaller integer
            }            
        }
    }

    return results;
};

const answer = threeSum([-1, 0, 1, 2, -1, -4]);
console.log(answer);
