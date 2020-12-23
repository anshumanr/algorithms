//Original at:
//https://github.com/eschwabe/interview-practice/blob/master/coursera/algorithms-part1/union-find/SuccessorWithDelete.java

import edu.princeton.cs.algs4.QuickFindUF;

class Successor {

    private QuickFindUF uf;

    public Successor(int N) {
        uf = new QuickFindUF(N);
    }

    public void remove(int x) {
        uf.union(x, x + 1);
    }

    public int successor(int x) {
        return uf.find(x);
    }

    public static void main(String[] args) {
        int N = 50;
        Successor swd = new Successor(50);
        System.out.println(swd.successor(10));
        swd.remove(11);
        swd.remove(13);
        swd.remove(12);
        swd.remove(10);
        swd.remove(9);
        swd.remove(15);
        System.out.println(swd.successor(8));
        System.out.println(swd.successor(9));
        System.out.println(swd.successor(10));
    }
}
