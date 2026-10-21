#ifndef DLIST_H
#define DLIST_H

#include <string>
#include <iostream>
using namespace std;

struct DNode {
    string data;
    DNode* prev;
    DNode* next;
    DNode(const string& v);
};

struct DList {
    DNode* head;
    DNode* tail;
    int count;

    DList();
    ~DList();

    void pushHead(const string& v);
    void pushTail(const string& v);
    void pushAfter(const string& target, const string& v);
    void pushBefore(const string& target, const string& v);

    void popHead();
    void popTail();
    void popAfter(const string& target);
    void popBefore(const string& target);
    void removeByValue(const string& v);

    bool find(const string& v);
    string at(int index);
    int size();
    void clear();

    void printForward();
    void printBackward();
    void printDetails();

    void save(ostream& os);
    void load(istream& is);
};

#endif