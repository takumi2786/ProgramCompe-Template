package solver;

// import java.util.Array;
import org.junit.Test;
import static org.hamcrest.MatcherAssert.*;
import static org.hamcrest.CoreMatchers.*;

public class SolverTest {
    // private HelloWorld helloWorld = new HelloWorld();

    @Test
    public void testMain() {
        System.out.println("Hello World Test!");
    }

    // solver test
    private Algo algo = new Algo();
    @Test
    public void testQuickSort() {
        int arr[] = {10,40, 29, 35, 1};
        int ans[] = algo.quickSort(arr);
        int correct[] = {1, 10, 29, 35, 40};
        assertThat(ans, is(correct));
    }
}
