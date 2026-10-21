#include "stack.h"

StackNode::StackNode(const string& v) : data(v), next(nullptr) {}

Stack::Stack() : topNode(nullptr), count(0) {}
Stack::~Stack() { clear(); }

void Stack::push(const string& v) {
    StackNode* n = new StackNode(v);
    n->next = topNode;
    topNode = n;
    count++;
}

string Stack::pop() {
    if (!topNode) return "";
    string v = topNode->data;
    StackNode* t = topNode;
    topNode = topNode->next;
    delete t;
    count--;
    return v;
}

string Stack::top() {
    return topNode ? topNode->data : "";
}

bool Stack::isEmpty() { return topNode == nullptr; }
int Stack::size() { return count; }

void Stack::clear() {
    while (topNode) { StackNode* t = topNode; topNode = topNode->next; delete t; }
    count = 0;
}

void Stack::printRec(StackNode* n) {
    if (!n) return;
    printRec(n->next);
    cout << n->data;
    if (n != topNode) cout << ", ";
}

void Stack::print() {
    cout << "Стек:    [";
    printRec(topNode);
    cout << "]  (вершина = " << top() << ")\n";
}

void Stack::printDetails() {
    cout << "Подробный вывод стека:\n";
    int pos = 0;
    for (StackNode* p = topNode; p; p = p->next) {
        cout << "  поз=" << pos++ << " addr=" << p
             << " data=\"" << p->data
             << "\" next=" << p->next << "\n";
    }
}

void Stack::saveRec(ostream& os, StackNode* n) {
    if (!n) return;
    saveRec(os, n->next);
    os << " " << n->data;
}

void Stack::save(ostream& os) { saveRec(os, topNode); }

void Stack::load(istream& is) {
    string v;
    while (is >> v) push(v);
}