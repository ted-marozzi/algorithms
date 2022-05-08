#include <string>
#include <map>
#include <deque>
#include <queue>
#include <stack>
#include <sstream>
#include <iostream>
#include <iomanip>
#include <cstdio>
#include <cmath>
#include <cstdlib>
#include <ctime>
#include <algorithm>
#include <vector>
#include <set>
#include <complex>
#include <list>

// #include "sort/sort.hpp"
// #include "util/util.hpp"

std::string solve(std::string n)
{
    return n;
}

int main()
{
    std::vector<std::string> input;
    for (std::string line; std::getline(std::cin, line);)
    {
        input.push_back(line);
    }

    int numTestCases = std::stoi(input[0]);

    for (auto i = 1; i <= numTestCases; i++)
    {
        auto result = solve(input[i]);
        std::cout << "Case #" << i << ": " << result << "\n";
    }

    return 0;
}
