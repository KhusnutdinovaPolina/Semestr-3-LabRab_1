#include "queue.h"

QueueNode::QueueNode(const string& v) : data(v), next(nullptr) {}

Queue::Queue() : frontNode(nullptr), rearNode(nullptr), count(0) {}
Queue::~Queue() { clear(); }

void Queue::push(const string& v) {
    QueueNode* n = new QueueNode(v);
    if (!rearNode) frontNode = rearNode = n;
    else { rearNode->next = n; rearNode = n; }
    count++;
}

string Queue::pop() {
    if (!frontNode) return "";
    string v = frontNode->data;
    QueueNode* t = frontNode;
    frontNode = frontNode->next;
    if (!frontNode) rearNode = nullptr;
    delete t;
    count--;
    return v;
}

string Queue::front() {
    return frontNode ? frontNode->data : "";
}

bool Queue::isEmpty() { return frontNode == nullptr; }
int Queue::size() { return count; }

void Queue::clear() {
    while (frontNode) { QueueNode* t = frontNode; frontNode = frontNode->next; delete t; }
    rearNode = nullptr;
    count = 0;
}

void Queue::print() {
    cout << "Очередь: [";
    for (QueueNode* p = frontNode; p; p = p->next) {
        cout << p->data;
        if (p->next) cout << " -> ";
    }
    cout << "]  (начало = " << front() << ")\n";
}

void Queue::printDetails() {
    cout << "Подробный вывод очереди:\n";
    int pos = 0;
    for (QueueNode* p = frontNode; p; p = p->next) {
        cout << "  поз=" << pos++ << " addr=" << p
             << " data=\"" << p->data
             << "\" next=" << p->next << "\n";
    }
}

void Queue::save(ostream& os) {
    for (QueueNode* p = frontNode; p; p = p->next) os << " " << p->data;
}

void Queue::load(istream& is) {
    string v;
    while (is >> v) push(v);
}