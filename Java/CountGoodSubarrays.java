import java.util.HashMap;
import java.util.Map;

public class CountGoodSubarrays {
    public static long countGood(int[] nums, int k) {
        Map<Integer, Integer> count = new HashMap<>();
        int pairs = 0, left = 0;
        long result = 0;

        for (int right = 0; right < nums.length; right++) {
            int num = nums[right];
            pairs += count.getOrDefault(num, 0);
            count.put(num, count.getOrDefault(num, 0) + 1);

            while (pairs >= k) {
                result += nums.length - right;
                count.put(nums[left], count.get(nums[left]) - 1);
                pairs -= count.get(nums[left]);
                left++;
            }
        }

        return result;
    }

    public static void main(String[] args) {
        int[] nums = {1, 1, 1, 1, 1};
        int k = 2;
        System.out.println(countGood(nums, k));
    }
}
