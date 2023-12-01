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
    int score2 = 0;
    int p1 = 0;
    int p2 = 0;
    string line;

    while (getline(infile, line)) {
        p1 = line[0] - 'A';
        p2 = line[2] - 'X';

        // Part 1
        score += p2 + 1 + ((p2 - p1 + 4) % 3) * 3;

        // Part 2
        score2 += p2 * 3 + (p1 + p2 + 2) % 3 + 1;
    }

    cout << "Day 2" << endl;
    cout << "The total score is: " << score << endl;
    cout << "The total correct score is: " << score2 << endl;
}
