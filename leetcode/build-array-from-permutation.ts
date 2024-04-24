function buildArray(nums: number[]): number[] {
    const arr = [];
    for (let i = 0; i < nums.length; i++) {
        let index = nums[i];
        arr.push(nums[index]);
    }
    return arr;
};

buildArray([5, 0, 1, 2, 3, 4]);      // [4,5,0,1,2,3]
// buildArray([0,2,1,5,3,4]);   // [0,1,2,4,5,3]


const buildArraySecondSolution = function(nums: number[]): number[] {
    return nums.map(a => nums[a]);
};