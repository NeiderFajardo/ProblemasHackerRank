#include <bits/stdc++.h>

using namespace std;

// Complete the commonChild function below.
int commonChild(string s1, string s2) {
	vector < vector<int> > r(s1.length()+1,vector<int>(s2.length()+1));
	for (int i=1;i<=s1.length();i++) {
		for (int j=1;j<=s2.length();j++) {        
			if (s1[i-1]==s2[j-1]){
				 r[i][j]=r[i-1][j-1]+1;
			 }
			else{
				r[i][j]=max(r[i][j-1], r[i-1][j]);
			}
    }
} 

	return r[s1.length()][s2.length()];
}

int main()
{
    string s1;
    getline(cin, s1);

    string s2;
    getline(cin, s2);

    int result = commonChild(s1, s2);

    cout<<result<<endl;

    return 0;
}

