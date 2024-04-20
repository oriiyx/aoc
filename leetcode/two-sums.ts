function twoSum(nums: number[], target: number): number[] {
    const map = new Map();

    for (let i = 0; i < nums.length; i++) {
        let deducted = target - nums[i];

        if (map.has(deducted)) {
            return [map.get(deducted), i];
        }

        map.set(nums[i], i);
    }

    return [];
};






console.log(twoSum([3, 2, 4], 6))

// twoSum([3,2,4], 6);