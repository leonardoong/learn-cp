#include<bits/stdc++.h>

using namespace std;

int main() {
    string a, b;
    cin >> a >> b;

    transform(a.begin(), a.end(), a.begin(), ::tolower);
    transform(b.begin(), b.end(), b.begin(), ::tolower);

    if (string(a) < string(b)) {
        cout << "-1";
    } else if ( string(b) < string(a)) {
        cout << "1";
    } else {
        cout << "0";
    }
}