function maxArea(height: number[]): number {
    let maxArea = 0;
    let left = 0;
    let right = height.length - 1;

    while (left < right) {
        // Calculate the area with the current pair of pointers
        const minHeight = Math.min(height[left], height[right]);
        const width = right - left;
        const currentArea = minHeight * width;

        // Update the maximum area found so far
        maxArea = Math.max(maxArea, currentArea);

        // Move the pointer pointing to the smaller number inside 
        /* 
            height[left] 1
            height[right] 7
            ==
            height[left] 8
            height[right] 7
            ==
            height[left] 8
            height[right] 3
            ==
            height[left] 8
            height[right] 8
            ==
            height[left] 8
            height[right] 4
            ==
            height[left] 8
            height[right] 5
            ==
            height[left] 8
            height[right] 2
            ==
            height[left] 8
            height[right] 6
        */
        if (height[left] < height[right]) {
            left++;
        } else {
            right--;
        }
    }

    return maxArea;
}

console.log(maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7]));
