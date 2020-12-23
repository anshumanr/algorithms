public class SoConnect
{
    private int[] id;
    private int[] sz;
    private int[] max;
    private int trees;
    private int members;

    public SoConnect(int N)
    {
        trees = N;
        members = N;
        id = new int[N];
        sz = new int[N];
        for (int i = 0; i < N; i++) id[i] = i;
        for (int i = 0; i < N; i++) id[i] = i;
        for (int i = 0; i < N; i++) max[i] = i;
    }

    private int root(int i)
    {
        while (i != id[i]) {
            id[i] = id[id[i]];
            i = id[i];
        }

        return i;
    }

    public boolean allconnected()
    {
        return trees == 1;
    }

    public boolean connected(int p, int q)
    {
        return root(p) == root(q);
    }

    public int find(int p) {
        int r = root(p);
        int max = p;

        for (int i = 0; i < members; i++) {
            if (id[i] > max) {
                max = id[i];
            }
        }

        return max;
    }

    public void union(int p, int q)
    {
        int i = root(p);
        int j = root(q);
        if (i == j)
            return;
        if (sz[i] >= sz[j]) {
            id[j] = i;
            sz[i] += sz[j];
            max[i] = Math.max(max[i], max[j]);
        }
        else {
            id[i] = j;
            sz[j] += sz[i];
            max[j] = Math.max(max[i], max[j]);
        }

        trees--;
    }

    public static void main(String[] args) {
        //feed input

        //create friendship (union) - by timestamp
    }
}
