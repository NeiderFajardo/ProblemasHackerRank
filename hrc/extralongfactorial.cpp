#include <bits/stdc++.h>

using namespace std;

// Complete the extraLongFactorials function below.
void extraLongFactorials(int n) {
    int respuesta = n;
    for (int i=0; i< n/2; i++){
        respuesta = respuesta*(n-i-1)*(i+1);
        cout<<n-i-1<<" "<<i+1<<endl;
    }

    cout<<respuesta<<endl;
}

int main()
{
    int n;
    cin >> n;
    cin.ignore(numeric_limits<streamsize>::max(), '\n');

    extraLongFactorials(n);

    return 0;
}
