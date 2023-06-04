#include <cstdlib>
#include <fstream>
#include <iostream>
#include <string>

using namespace std;

int main() {
    ifstream infile;

    infile.open("../day2/day2.in");

    if (!infile) {
        cerr << "Unable to open the file!";
        exit(1);
    }

    int score = 0;
    int p1 = 0;
    int p2 = 0;
    string line;

    while (getline(infile, line)) {
        p1 = line[0] - 'A' + 1;
        p2 = line[2] - 'X' + 1;
        score += p1;
        if (p1 == 1 && p2 == 3) {
            score += 6;
        } else if (p1 == 2  && p2 == 1) {
            score += 6;
        } else if (p1 == 3 && p2 == 2) {
            score += 6;
        } else if (p1 == p2) {
            score += 3;
        }
    }

    cout << "Day 2" << endl;
    cout << "The total score is: " << score << endl;
}
