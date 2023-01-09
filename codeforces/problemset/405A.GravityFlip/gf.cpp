#include <bits/stdc++.h>

using namespace std;

int main()
{
    int a;
    cin >> a;

    int x[a] = {};
    for (int i = 0; i < a; i++)
    {
        int temp;
        cin >> temp;
        x[i] = temp;
    }

    sort(x, x + a);

    for (int i = 0; i < a; i++)
    {
        cout << x[i] << " ";
    }

    return 0;
}