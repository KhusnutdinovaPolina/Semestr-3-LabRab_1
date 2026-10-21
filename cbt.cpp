#include "cbt.h"

TNode::TNode(const string& v) : data(v), left(nullptr), right(nullptr) {}

CBT::CBT() : root(nullptr), count(0) {}
CBT::~CBT() { clearRec(root); }

TNode* CBT::nodeByIndex(TNode* n, int cur, int target) {
    if (!n) return nullptr;
    if (cur == target) return n;
    TNode* l = nodeByIndex(n->left, 2*cur + 1, target);
    if (l) return l;
    return nodeByIndex(n->right, 2*cur + 2, target);
}

int CBT::height(TNode* n) {
    if (!n) return 0;
    int l = height(n->left);
    int r = height(n->right);
    return 1 + (l > r ? l : r);
}

void CBT::insert(const string& v) {
    TNode* n = new TNode(v);
    if (count == 0) { root = n; count = 1; return; }
    int idx = count;
    int parentIdx = (idx - 1) / 2;
    bool isLeft = (idx % 2 == 1);
    TNode* parent = nodeByIndex(root, 0, parentIdx);
    if (isLeft) parent->left = n;
    else        parent->right = n;
    count++;
}

TNode* CBT::findNode(TNode* n, const string& v) {
    if (!n) return nullptr;
    if (n->data == v) return n;
    TNode* l = findNode(n->left, v);
    if (l) return l;
    return findNode(n->right, v);
}

bool CBT::find(const string& v) { return findNode(root, v) != nullptr; }

bool CBT::isComplete() {
    if (!root) return true;
    
    TNode* queue[1000];
    int front = 0, rear = 0;

    queue[rear++] = root;
    bool seenNull = false;

    while (front < rear) {
        TNode* cur = queue[front++];

        if (cur) {
            if (seenNull) return false;

            queue[rear++] = cur->left;
            queue[rear++] = cur->right;
        } else {
            seenNull = true;
        }
    }
    return true;
}

int CBT::size() { return count; }

void CBT::clearRec(TNode* n) {
    if (!n) return;
    clearRec(n->left);
    clearRec(n->right);
    delete n;
}

void CBT::clear() { clearRec(root); root = nullptr; count = 0; }

void CBT::print() {
    cout << "Уровни:   [";
    for (int i = 0; i < count; i++) {
        TNode* t = nodeByIndex(root, 0, i);
        if (i) cout << ", ";
        cout << t->data;
    }
    cout << "]\n";
    cout << "Complete: " << (isComplete() ? "ДА" : "НЕТ") << "\n";

    if (!root) { cout << "(дерево пусто)\n"; return; }
    cout << "Дерево:\n";

    int h = height(root);
    int maxWidth = (1 << h) - 1;

    int cell = 3;
    for (int i = 0; i < count; i++) {
        TNode* t = nodeByIndex(root, 0, i);
        if ((int)t->data.size() + 2 > cell) cell = (int)t->data.size() + 2;
    }
    int totalWidth = maxWidth * cell;

    for (int level = 0; level < h; level++) {
        string line(totalWidth, ' ');
        int nodesOnLevel = 1 << level;
        int spanPerNode = totalWidth / nodesOnLevel;

        for (int j = 0; j < nodesOnLevel; j++) {
            int idx = (1 << level) - 1 + j;
            if (idx >= count) break;
            TNode* t = nodeByIndex(root, 0, idx);

            int center = j * spanPerNode + spanPerNode / 2;
            int start  = center - (int)t->data.size() / 2;
            if (start < 0) start = 0;

            for (size_t k = 0; k < t->data.size() && start + (int)k < totalWidth; k++)
                line[start + k] = t->data[k];
        }
        while (!line.empty() && line.back() == ' ') line.pop_back();
        cout << line << "\n";

        if (level + 1 < h) {
            string branches(totalWidth, ' ');
            for (int j = 0; j < nodesOnLevel; j++) {
                int idx = (1 << level) - 1 + j;
                if (idx >= count) break;
                TNode* t = nodeByIndex(root, 0, idx);
                if (!t->left && !t->right) continue;

                int center = j * spanPerNode + spanPerNode / 2;

                if (t->left) {
                    int lc = center - spanPerNode / 2;
                    for (int k = lc; k < center; k++)
                        if (k >= 0 && k < totalWidth) branches[k] = '-';
                    if (lc >= 0 && lc < totalWidth) branches[lc] = '+';
                }
                if (t->right) {
                    int rc = center + spanPerNode / 2;
                    for (int k = center + 1; k <= rc; k++)
                        if (k >= 0 && k < totalWidth) branches[k] = '-';
                    if (rc >= 0 && rc < totalWidth) branches[rc] = '+';
                }
                if (center >= 0 && center < totalWidth) branches[center] = '+';
            }
            while (!branches.empty() && branches.back() == ' ') branches.pop_back();
            cout << branches << "\n";
        }
    }
}

void CBT::printDetails() {
    cout << "Подробный вывод дерева:\n";
    printDetailsRec(root, 0);
}

void CBT::printDetailsRec(TNode* n, int level) {
    if (!n) return;
    cout << "  уровень " << level
         << " addr=" << n
         << " data=\"" << n->data << "\""
         << " left=" << n->left
         << " right=" << n->right << "\n";
    printDetailsRec(n->left, level + 1);
    printDetailsRec(n->right, level + 1);
}

void CBT::save(ostream& os) {
    for (int i = 0; i < count; i++) {
        TNode* t = nodeByIndex(root, 0, i);
        os << " " << t->data;
    }
}

void CBT::load(istream& is) {
    string v;
    while (is >> v) insert(v);
}