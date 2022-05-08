
export module helloworld;
import<iostream>;

export void hello()
{ // export declaration
    std::cout << "Hello world!\n";
}

//     "command": "g++ -c -fmodules-ts -x c++-system-header -std=c++20 iostream string vector && g++ -std=c++20 -fmodules-ts sort/helloworld.cpp main.cpp -o main.out"
