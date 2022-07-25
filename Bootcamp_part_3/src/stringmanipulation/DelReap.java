package stringmanipulation;
import java.util.*;
public class DelReap {
	
	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Scanner sc = new Scanner(System.in);
		System.out.println("Enter a string");
		String str = sc.next();
		
		for (int i=0;i<str.length();i++) {
			char c=str.charAt(i);
			if (str.indexOf(c)<i)
				continue;
			else
				System.out.print(c);
			
		}
		
         sc.close();
	}

}
