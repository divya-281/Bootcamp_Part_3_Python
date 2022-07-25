package arraymanipulation;
import java.util.*;
public class MaxMin {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Scanner sc =new Scanner(System.in);
		System.out.println("Enter Size of the array");
		int n=sc.nextInt();
		int a[]=new int[n];
		System.out.println("Enter "+n+" elements");
		for (int i=0;i<n;i++) {
			a[i]=sc.nextInt();
		}
		int max=a[0],min=a[0];
		for (int i=0;i<n;i++) {
			if (a[i]>max) {
				max=a[i];
			}
			if (a[i]<min) {
				min=a[i];
			}
			
		}
		System.out.println("Largest : "+max);
		System.out.println("Smallest : "+min);
		sc.close();

	}

}
