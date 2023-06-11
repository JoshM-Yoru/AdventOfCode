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
    bool tracker[100];

    while (getline(infile, line)) {

        for (int i = 0; i < line.length() / 2; i++) {
           if (line[i] < 'a') {
               tracker[line[i] - 'A' + 26] = true;
           } else {
               tracker[line[i] - 'A' + 26] = true;
               tracker[line[i] - 'a'] = true;
           }
        }
        for (int i = line.length() / 2; i < line.length(); i++) {
           if (line[i] < 'a' && tracker[line[i] - 'A' + 26] == true) {
               total += line[i] - 'A' + 27;
               break;
           } else if (line[i] >= 'a' && tracker[line[i] - 'a'] == true){
               total += line[i] - 'a' + 1;
               break;
           }
        }
    }

    cout << "Day 3" << endl;
    cout << "The total is: " << total << endl;
    cout << "" << endl;
}
