package stringmanipulation;
import java.util.*;
public class RemChar {
	public static void main(String args[]) {
		Scanner sc = new Scanner(System.in);
		System.out.println("Enter a string");
		String str = sc.next();
		String new_str="";
		int flag=0;
		System.out.println("Enter the character to be removed");
		char ch= sc.next().charAt(0);
		for(int i=0; i<str.length();i++)
		{
			char c= str.charAt(i);
			if (c==ch) {
				flag=1;
				continue;
			}
			else {
				new_str+=c;
				
			}
		}
		if (flag==1) {
			System.out.println("New String is "+new_str);
		}
		else {
			System.out.println(ch +" not found in "+str);
		}
		
		sc.close();
	}

}
