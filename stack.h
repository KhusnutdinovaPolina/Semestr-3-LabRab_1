#ifndef STACK_H
#define STACK_H

#include <string>
#include <iostream>
using namespace std;

struct StackNode {
    string data;
    StackNode* next;
    StackNode(const string& v);
};

struct Stack {
    StackNode* topNode;
    int count;

    Stack();
    ~Stack();

    void push(const string& v);
    string pop();
    string top();
    bool isEmpty();
    int size();
    void clear();

    void print();
    void printRec(StackNode* n);
    void printDetails();

    void save(ostream& os);
    void saveRec(ostream& os, StackNode* n);
    void load(istream& is);
};

#endif