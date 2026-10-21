#ifndef SLIST_H
#define SLIST_H

#include <string>
#include <iostream>
using namespace std;

struct SNode {
    string data;
    SNode* next;
    SNode(const string& v);
};

struct SList {
    SNode* head;
    int count;

    SList();
    ~SList();

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
    void printReverse();
    void printReverseRec(SNode* n);
    void printDetails();

    void save(ostream& os);
    void load(istream& is);
};

#endif