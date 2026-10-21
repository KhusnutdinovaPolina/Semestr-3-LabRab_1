#ifndef QUEUE_H
#define QUEUE_H

#include <string>
#include <iostream>
using namespace std;

struct QueueNode {
    string data;
    QueueNode* next;
    QueueNode(const string& v);
};

struct Queue {
    QueueNode* frontNode;
    QueueNode* rearNode;
    int count;

    Queue();
    ~Queue();

    void push(const string& v);
    string pop();
    string front();
    bool isEmpty();
    int size();
    void clear();

    void print();
    void printDetails();

    void save(ostream& os);
    void load(istream& is);
};

#endif