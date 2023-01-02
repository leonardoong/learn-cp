#include <bits/stdc++.h>

using namespace std;

int main() {
    string a;
    cin >> a;

    int c = 0;

    for (int i = 0; i < a.length(); i++){
        if (a[i - 1] != a[i]) {
            c = 0;
        }
        c++;
        if (c == 7) {
            cout << "YES";
            return 0;
        }
    }

    cout << "NO";
    return 0;
}