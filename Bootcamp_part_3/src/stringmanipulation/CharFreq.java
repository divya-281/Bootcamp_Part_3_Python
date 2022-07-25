package stringmanipulation;
import java.util.*;
public class CharFreq {
	

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Scanner sc = new Scanner(System.in);
		System.out.println("Enter a string");
		String str = sc.next();
		char c,ch;
		for (int i=0;i<str.length();i++)
		{
			
			int count=0;
			
			ch=str.charAt(i);
			if (ch!=' ')
			{
			
				for (int j=0;j<str.length();j++) 
				{
					c=str.charAt(j);
					if (c==ch)
					{
						count++;
					}
					
				}
				System.out.println(ch+" occurs "+count +" number of times");
				str= str.replace(ch, ' ');
			}
		}
		sc.close();

	}

}
