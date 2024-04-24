function getConcatenation(nums: number[]): number[] {
    const fixedLength = nums.length;
    let movingLenth = nums.length;

    for (let i = 0; i < fixedLength; i++) {
        nums[movingLenth] = nums[i];
        movingLenth++;
    }

    return nums;
};

getConcatenation([1, 3, 2, 1]);