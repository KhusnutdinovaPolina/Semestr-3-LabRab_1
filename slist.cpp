#include "slist.h"

SNode::SNode(const string& v) : data(v), next(nullptr) {}

SList::SList() : head(nullptr), count(0) {}
SList::~SList() { clear(); }

void SList::pushHead(const string& v) {
    SNode* n = new SNode(v);
    n->next = head;
    head = n;
    count++;
}

void SList::pushTail(const string& v) {
    SNode* n = new SNode(v);
    if (!head) { head = n; count++; return; }
    SNode* p = head;
    while (p->next) p = p->next;
    p->next = n;
    count++;
}

void SList::pushAfter(const string& target, const string& v) {
    SNode* p = head;
    while (p && p->data != target) p = p->next;
    if (!p) return;
    SNode* n = new SNode(v);
    n->next = p->next;
    p->next = n;
    count++;
}

void SList::pushBefore(const string& target, const string& v) {
    if (!head) return;
    if (head->data == target) { pushHead(v); return; }
    SNode* p = head;
    while (p->next && p->next->data != target) p = p->next;
    if (!p->next) return;
    SNode* n = new SNode(v);
    n->next = p->next;
    p->next = n;
    count++;
}

void SList::popHead() {
    if (!head) return;
    SNode* t = head;
    head = head->next;
    delete t;
    count--;
}

void SList::popTail() {
    if (!head) return;
    if (!head->next) { delete head; head = nullptr; count--; return; }
    SNode* p = head;
    while (p->next->next) p = p->next;
    delete p->next;
    p->next = nullptr;
    count--;
}

void SList::popAfter(const string& target) {
    SNode* p = head;
    while (p && p->data != target) p = p->next;
    if (!p || !p->next) return;
    SNode* t = p->next;
    p->next = t->next;
    delete t;
    count--;
}

void SList::popBefore(const string& target) {
    if (!head || head->data == target) return;
    if (head->next && head->next->data == target) { popHead(); return; }
    SNode* p = head;
    while (p->next && p->next->next && p->next->next->data != target) p = p->next;
    if (!p->next || !p->next->next) return;
    SNode* t = p->next;
    p->next = t->next;
    delete t;
    count--;
}

void SList::removeByValue(const string& v) {
    if (!head) return;
    if (head->data == v) { popHead(); return; }
    SNode* p = head;
    while (p->next && p->next->data != v) p = p->next;
    if (p->next) {
        SNode* t = p->next;
        p->next = t->next;
        delete t;
        count--;
    }
}

bool SList::find(const string& v) {
    for (SNode* p = head; p; p = p->next)
        if (p->data == v) return true;
    return false;
}

string SList::at(int index) {
    if (index < 0 || index >= count) return "";
    SNode* p = head;
    for (int i = 0; i < index; i++) p = p->next;
    return p->data;
}

int SList::size() { return count; }

void SList::clear() {
    while (head) { SNode* t = head; head = head->next; delete t; }
    count = 0;
}

void SList::printForward() {
    cout << "Список:  [";
    for (SNode* p = head; p; p = p->next) {
        cout << p->data;
        if (p->next) cout << " -> ";
    }
    cout << "]\n";
}

void SList::printReverseRec(SNode* n) {
    if (!n) return;
    printReverseRec(n->next);
    cout << n->data;
    if (n != head) cout << " <- ";
}

void SList::printReverse() {
    cout << "Назад:   [";
    printReverseRec(head);
    cout << "]\n";
}

void SList::printDetails() {
    cout << "Подробный вывод списка:\n";
    int i = 0;
    for (SNode* p = head; p; p = p->next) {
        cout << "  [" << i++ << "] addr=" << p
             << " data=\"" << p->data
             << "\" next=" << p->next << "\n";
    }
}

void SList::save(ostream& os) {
    for (SNode* p = head; p; p = p->next) os << " " << p->data;
}

void SList::load(istream& is) {
    string v;
    while (is >> v) pushTail(v);
}