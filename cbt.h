#ifndef CBT_H
#define CBT_H

#include <string>
#include <iostream>
using namespace std;

struct TNode {
    string data;
    TNode* left;
    TNode* right;
    TNode(const string& v);
};

struct CBT {
    TNode* root;
    int count;

    CBT();
    ~CBT();

    void insert(const string& v);
    bool find(const string& v);
    bool isComplete();
    int size();
    void clear();

    TNode* nodeByIndex(TNode* n, int cur, int target);
    TNode* findNode(TNode* n, const string& v);
    int height(TNode* n);
    void clearRec(TNode* n);

    void print();
    void printDetails();
    void printDetailsRec(TNode* n, int level);

    void save(ostream& os);
    void load(istream& is);
};

#endif