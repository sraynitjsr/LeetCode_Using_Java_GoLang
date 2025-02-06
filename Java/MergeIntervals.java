import java.util.*;

public class MergeIntervals {
    public static void main(String[] args) {
        List<List<Integer>> intervals = new ArrayList<>();
        intervals.add(Arrays.asList(1, 3));
        intervals.add(Arrays.asList(2, 6));
        intervals.add(Arrays.asList(8, 10));
        intervals.add(Arrays.asList(15, 18));

        Collections.sort(intervals, Comparator.comparingInt(a -> a.get(0)));

        List<List<Integer>> merged = new ArrayList<>();
        merged.add(intervals.get(0));

        for (int i = 1; i < intervals.size(); i++) {
            List<Integer> last = merged.get(merged.size() - 1);
            List<Integer> current = intervals.get(i);

            if (current.get(0) <= last.get(1)) {
                last.set(1, Math.max(last.get(1), current.get(1)));
            } else {
                merged.add(current);
            }
        }

        System.out.println(merged);
    }
}
