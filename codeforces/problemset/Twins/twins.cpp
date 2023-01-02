#include <bits/stdc++.h>
#include <algorithm>

using namespace std;

bool compare (int a, int b) {
    if (a > b) return true;
    return false;
}

int main() {
    int x;
    cin >> x;

    int ar[105];
    int mySum = 0;
    int broSum = 0;
    int counter = 0;

    for (int i = 0; i < x; i++){
        cin >> ar[i];
        broSum += ar[i];
    }

    sort(ar, ar + x, compare);

    
    for (int i = 0; i < x; i++) {
        mySum += ar[i];
        broSum -= ar[i];
        counter++;
        if (mySum > broSum) {
            break;
        }
    }

    cout << counter << endl;

    return 0;
}