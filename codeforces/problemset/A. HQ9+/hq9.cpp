#include <bits/stdc++.h>

using namespace std;

int main()
{
    string a;
    cin >> a;

    for (int i = 0; i < a.length(); i++)
    {
        if (a[i] == 72 || a[i] == 81 || a[i] == 57)
        {
            cout << "YES";
            return 0;
        }
    }

    cout << "NO";
    return 0;
}