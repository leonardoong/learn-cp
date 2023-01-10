#include <bits/stdc++.h>

using namespace std;

int main()
{
    string a;
    cin >> a;

    string r = "";
    for (int i = 0; i < a.length(); i++)
    {
        a[i] = tolower(a[i]);
        if (a[i] == 'a' || a[i] == 'e' || a[i] == 'i' || a[i] == 'o' || a[i] == 'u' || a[i] == 'y')
        {
            continue;
        }

        r.append(".");
        string s(1, a[i]);
        r.append(s);
    }

    cout << r;
    return 0;
}