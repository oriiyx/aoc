function findMedianSortedArrays(nums1: number[], nums2:number[]): number {
    let i = 0, j = 0;
    let merged = [];

    // Merge the two sorted arrays
    while (i < nums1.length && j < nums2.length) {
        if (nums1[i] < nums2[j]) {
            merged.push(nums1[i++]);
        } else {
            merged.push(nums2[j++]);
        }
    }

    // If there are remaining elements in nums1
    while (i < nums1.length) {
        merged.push(nums1[i++]);
    }

    // If there are remaining elements in nums2
    while (j < nums2.length) {
        merged.push(nums2[j++]);
    }

    let total = merged.length;
    let half = Math.floor(total / 2);

    if (total % 2 === 0) {
        // If even, average of two middle numbers
        return (merged[half - 1] + merged[half]) / 2;
    } else {
        // If odd, middle number
        return merged[half];
    }
}

console.log(findMedianSortedArrays([1,3], [2]));