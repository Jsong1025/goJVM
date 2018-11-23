public class Test1 {
	public static int staticVar;
	public int instanceVar;

	public static void main(String[] args) {
		int x = 32768;
		Test1 test = new Test1();
		Test1.staticVar = x;
		x = Test1.staticVar;
		test.instanceVar = x;
		x = test.instanceVar;
		Object obj = test;
		if(obj instanceof Test1) {
			test = (Test1)obj;
			System.out.println(test.instanceVar);
		}
	}
}
