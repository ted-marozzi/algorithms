#pragma once

#include <vector>
#include "../util/util.hpp"

template <typename T>
void simpleSort(std::vector<T> &vector)
{
    for (auto i = 0; i < vector.size(); i++)
        for (auto j = 0; j < vector.size(); j++)
            if (vector[i] < vector[j])
                swap(vector[i], vector[j]);
}