package stringmanipulation;
import java.util.*;
public class Palindrome {
	public static void main (String args[]) {
		Scanner sc = new Scanner(System.in);
		System.out.println("Enter a string");
		String s=sc.next();
		int l=s.length();
		String rev="";
		String s1=s.toLowerCase();
		for (int i=0;i<l;i++) {
			char c= s1.charAt(i);
			rev=c+rev;
			}
		if (s1.equals(rev)) {
			System.out.println(s+" is a palindrome");
		}
		else {
			System.out.println(s+" is not a palindrome");
		}
		sc.close();
	}

}
