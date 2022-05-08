#pragma once
#include <iostream>
#include <vector>
#include "../util/util.hpp"

template <typename T>
void printVector(const std::vector<T> &vector)
{
    std::cout << "{";
    for (int value : vector)
        std::cout << " " << value;
    std::cout << " }\n";
}

template <typename T>
void swap(T &a, T &b)
{
    auto temp = a;
    a = b;
    b = temp;
}
