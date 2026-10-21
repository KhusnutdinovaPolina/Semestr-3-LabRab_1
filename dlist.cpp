#include "dlist.h"

DNode::DNode(const string& v) : data(v), prev(nullptr), next(nullptr) {}

DList::DList() : head(nullptr), tail(nullptr), count(0) {}
DList::~DList() { clear(); }

void DList::pushHead(const string& v) {
    DNode* n = new DNode(v);
    n->next = head;
    if (head) head->prev = n;
    head = n;
    if (!tail) tail = n;
    count++;
}

void DList::pushTail(const string& v) {
    DNode* n = new DNode(v);
    n->prev = tail;
    if (tail) tail->next = n;
    tail = n;
    if (!head) head = n;
    count++;
}

void DList::pushAfter(const string& target, const string& v) {
    DNode* p = head;
    while (p && p->data != target) p = p->next;
    if (!p) return;
    DNode* n = new DNode(v);
    n->next = p->next; n->prev = p;
    if (p->next) p->next->prev = n; else tail = n;
    p->next = n;
    count++;
}

void DList::pushBefore(const string& target, const string& v) {
    DNode* p = head;
    while (p && p->data != target) p = p->next;
    if (!p) return;
    DNode* n = new DNode(v);
    n->prev = p->prev; n->next = p;
    if (p->prev) p->prev->next = n; else head = n;
    p->prev = n;
    count++;
}

void DList::popHead() {
    if (!head) return;
    DNode* t = head;
    head = head->next;
    if (head) head->prev = nullptr; else tail = nullptr;
    delete t;
    count--;
}

void DList::popTail() {
    if (!tail) return;
    DNode* t = tail;
    tail = tail->prev;
    if (tail) tail->next = nullptr; else head = nullptr;
    delete t;
    count--;
}

void DList::popAfter(const string& target) {
    DNode* p = head;
    while (p && p->data != target) p = p->next;
    if (!p || !p->next) return;
    DNode* t = p->next;
    p->next = t->next;
    if (t->next) t->next->prev = p; else tail = p;
    delete t;
    count--;
}

void DList::popBefore(const string& target) {
    DNode* p = head;
    while (p && p->data != target) p = p->next;
    if (!p || !p->prev) return;
    DNode* t = p->prev;
    p->prev = t->prev;
    if (t->prev) t->prev->next = p; else head = p;
    delete t;
    count--;
}

void DList::removeByValue(const string& v) {
    DNode* p = head;
    while (p && p->data != v) p = p->next;
    if (!p) return;
    if (p->prev) p->prev->next = p->next; else head = p->next;
    if (p->next) p->next->prev = p->prev; else tail = p->prev;
    delete p;
    count--;
}

bool DList::find(const string& v) {
    for (DNode* p = head; p; p = p->next)
        if (p->data == v) return true;
    return false;
}

string DList::at(int index) {
    if (index < 0 || index >= count) return "";
    DNode* p = head;
    for (int i = 0; i < index; i++) p = p->next;
    return p->data;
}

int DList::size() { return count; }

void DList::clear() {
    while (head) { DNode* t = head; head = head->next; delete t; }
    tail = nullptr;
    count = 0;
}

void DList::printForward() {
    cout << "Вперёд:   [";
    for (DNode* p = head; p; p = p->next) {
        cout << p->data;
        if (p->next) cout << " <-> ";
    }
    cout << "]\n";
}

void DList::printBackward() {
    cout << "Назад:    [";
    for (DNode* p = tail; p; p = p->prev) {
        cout << p->data;
        if (p->prev) cout << " <-> ";
    }
    cout << "]\n";
}

void DList::printDetails() {
    cout << "Подробный вывод:\n";
    int i = 0;
    for (DNode* p = head; p; p = p->next) {
        cout << "  [" << i++ << "] addr=" << p
             << " prev=" << p->prev
             << " data=\"" << p->data
             << "\" next=" << p->next << "\n";
    }
}

void DList::save(ostream& os) {
    for (DNode* p = head; p; p = p->next) os << " " << p->data;
}

void DList::load(istream& is) {
    string v;
    while (is >> v) pushTail(v);
}