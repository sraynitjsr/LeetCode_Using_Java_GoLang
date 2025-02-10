public class MaxWaterContainer {
    public int maxArea(int[] height) {
        int left = 0, right = height.length - 1, maxWater = 0;
        while (left < right) {
            int h = Math.min(height[left], height[right]);
            maxWater = Math.max(maxWater, h * (right - left));
            if (height[left] < height[right]) {
                left++;
            } else {
                right--;
            }
        }
        return maxWater;
    }

    public static void main(String[] args) {
        MaxWaterContainer obj = new MaxWaterContainer();
        int[] height = {1, 8, 6, 2, 5, 4, 8, 3, 7};
        System.out.println(obj.maxArea(height));
    }
}
