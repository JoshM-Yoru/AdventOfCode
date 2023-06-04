#include <cstdlib>
#include <fstream>
#include <iostream>
#include <string>

using namespace std;

int main() {
    ifstream infile;

    infile.open("../day1/day1.in");

    if (!infile) {
        cerr << "Unable to open the file!";
        exit(1);
    }

    int top3[3];
    int curr = 0;
    string line;

    while (getline(infile, line)) {
        if (line.empty()) {
            if (top3[0] < curr) {
                top3[2] = top3[1];
                top3[1] = top3[0];
                top3[0] = curr;
            } else if (top3[1] < curr) {
                top3[2] = top3[1];
                top3[1] = curr;
            } else if (top3[2] < curr) {
                top3[2] = curr;
            }

            curr = 0;
            continue;
        }
        curr += stoi(line);
    }
    cout << "Day 1" << endl;
    cout << "The most calories are: " << top3[0] << endl;
    cout << "The total of the 3 most calories are: " << top3[0] + top3[1] + top3[2] << endl;
    cout << "" << endl;
}
