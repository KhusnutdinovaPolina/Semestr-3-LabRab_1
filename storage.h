#ifndef STORAGE_H
#define STORAGE_H

#include <map>
#include <string>
#include <fstream>
#include <sstream>
#include "array.h"
#include "slist.h"
#include "dlist.h"
#include "stack.h"
#include "queue.h"
#include "cbt.h"
using namespace std;

extern map<string, Array> arrays;
extern map<string, SList> slists;
extern map<string, DList> dlists;
extern map<string, Stack> stacks;
extern map<string, Queue> queues;
extern map<string, CBT>   trees;

inline void splitWords(const string& s, string* out, int& cnt) {
    cnt = 0;
    istringstream ss(s);
    string w;
    while (ss >> w && cnt < 100) out[cnt++] = w;
}

void saveFile(const string& filename);
void loadFile(const string& filename);

#endif