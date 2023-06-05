#include <cstdlib>
#include <fstream>
#include <iostream>
#include <string>

using namespace std;

int main() {
    ifstream infile;

    infile.open("../day3/day3.in");

    if (!infile) {
        cerr << "Unable to open the file!";
        exit(1);
    }

    string line;
    int total = 0;

    while (getline(infile, line)) {
        int tracker[52];

        for (int i = 0; i < line.length() / 2; i++) {
           if (line[i] >= 'A') {
               tracker[line[i] - 'A'] = 1;
           } else {
               tracker[line[i] - 'a' + 26] = 1;
           }
        }
        for (int i = line.length() / 2; i < line.length(); i++) {
           if (line[i] >= 'A' && tracker[line[i] - 'A'] == 1) {
               total += line[i] - 'A';
               tracker[line[i] - 'A'] = 2;
           } else if (line[i] < 'A' && tracker[line[i] - 'a' + 26] == 1){
               total += line[i] - 'a' + 26;
               tracker[line[i] - 'a' + 26] = 2;
           }
        }
    }

    cout << "Day 3" << endl;
    cout << "The total is: " << total << endl;
    cout << "" << endl;
}
