#ifndef ARRAY_H
#define ARRAY_H

#include <string>
#include <iostream>
using namespace std;

struct Array {
    string* data;
    int size;
    int capacity;

    Array();
    ~Array();

    void pushBack(const string& v);
    void insertAt(int index, const string& v);
    string get(int index);
    void set(int index, const string& v);
    void removeAt(int index);
    int length();
    void clear();

    void print();
    void printDetails();

    void save(ostream& os);
    void load(istream& is);
};

#endif