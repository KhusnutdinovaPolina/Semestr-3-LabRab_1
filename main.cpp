#include "storage.h"
#include <iostream>
using namespace std;

int main(int argc, char* argv[]) {
    string filename, query;

    for (int i = 1; i < argc; i++) {
        string a = argv[i];
        if (a == "--file"  && i+1 < argc) filename = argv[++i];
        else if (a == "--query" && i+1 < argc) query = argv[++i];
        else if (!a.empty() && a[0] != '-' && filename.empty()) filename = a;
    }

    if (filename.empty() || query.empty()) {
        cout << "Использование: ./dbms --file data.db --query 'КОМАНДА'\n";
        return 1;
    }

    loadFile(filename);

    string p[100]; int pc = 0;
    splitWords(query, p, pc);
    if (pc == 0) return 1;

    string c = p[0];
    bool changed = false;

    if (c == "MPUSH" && pc == 3) {
        arrays[p[1]].pushBack(p[2]); changed = true;
    }
    else if (c == "MPUSH" && pc == 4) {
        arrays[p[1]].insertAt(stoi(p[2]), p[3]); changed = true;
    }
    else if (c == "MGET" && pc == 3) {
        cout << arrays[p[1]].get(stoi(p[2])) << "\n";
    }
    else if (c == "MSET" && pc == 4) {
        arrays[p[1]].set(stoi(p[2]), p[3]); changed = true;
    }
    else if (c == "MDEL" && pc == 3) {
        arrays[p[1]].removeAt(stoi(p[2])); changed = true;
    }
    else if (c == "MLEN" && pc == 2) {
        cout << arrays[p[1]].length() << "\n";
    }
    else if (c == "FPUSH" && pc == 4) {
        auto& L = slists[p[1]];
        if      (p[2] == "head") L.pushHead(p[3]);
        else if (p[2] == "tail") L.pushTail(p[3]);
        changed = true;
    }
    else if (c == "FPUSH" && pc == 5) {
        auto& L = slists[p[1]];
        if      (p[2] == "after")  L.pushAfter (p[3], p[4]);
        else if (p[2] == "before") L.pushBefore(p[3], p[4]);
        changed = true;
    }
    else if (c == "FDEL" && pc == 3) {
        auto& L = slists[p[1]];
        if      (p[2] == "head") L.popHead();
        else if (p[2] == "tail") L.popTail();
        else                     L.removeByValue(p[2]);
        changed = true;
    }
    else if (c == "FDEL" && pc == 4) {
        auto& L = slists[p[1]];
        if      (p[2] == "after")  L.popAfter (p[3]);
        else if (p[2] == "before") L.popBefore(p[3]);
        changed = true;
    }
    else if (c == "FGET" && pc == 3) {
        auto& L = slists[p[1]];
        if (p[2] == "forward") L.printForward();
        else if (p[2] == "reverse") L.printReverse();
        else if (p[2] == "details") L.printDetails();
        else cout << (L.find(p[2]) ? "ИСТИНА" : "ЛОЖЬ") << "\n";
    }
    else if (c == "LPUSH" && pc == 4) {
        auto& L = dlists[p[1]];
        if      (p[2] == "head") L.pushHead(p[3]);
        else if (p[2] == "tail") L.pushTail(p[3]);
        changed = true;
    }
    else if (c == "LPUSH" && pc == 5) {
        auto& L = dlists[p[1]];
        if      (p[2] == "after")  L.pushAfter (p[3], p[4]);
        else if (p[2] == "before") L.pushBefore(p[3], p[4]);
        changed = true;
    }
    else if (c == "LDEL" && pc == 3) {
        auto& L = dlists[p[1]];
        if      (p[2] == "head") L.popHead();
        else if (p[2] == "tail") L.popTail();
        else                     L.removeByValue(p[2]);
        changed = true;
    }
    else if (c == "LDEL" && pc == 4) {
        auto& L = dlists[p[1]];
        if      (p[2] == "after")  L.popAfter (p[3]);
        else if (p[2] == "before") L.popBefore(p[3]);
        changed = true;
    }
    else if (c == "LGET" && pc == 3) {
        auto& L = dlists[p[1]];
        if      (p[2] == "forward")  L.printForward();
        else if (p[2] == "backward") L.printBackward();
        else if (p[2] == "details")  L.printDetails();
        else cout << (L.find(p[2]) ? "ИСТИНА" : "ЛОЖЬ") << "\n";
    }
    else if (c == "SPUSH" && pc == 3) { stacks[p[1]].push(p[2]); changed = true; }
    else if (c == "SPOP"  && pc == 2) { cout << stacks[p[1]].pop() << "\n"; changed = true; }
    else if (c == "SGET"  && pc == 2) { cout << stacks[p[1]].top() << "\n"; }
    else if (c == "QPUSH" && pc == 3) { queues[p[1]].push(p[2]); changed = true; }
    else if (c == "QPOP"  && pc == 2) { cout << queues[p[1]].pop() << "\n"; changed = true; }
    else if (c == "QGET"  && pc == 2) { cout << queues[p[1]].front() << "\n"; }
    else if (c == "TINSERT" && pc == 3) { trees[p[1]].insert(p[2]); changed = true; }
    else if (c == "TGET"    && pc == 3) {
        cout << (trees[p[1]].find(p[2]) ? "ИСТИНА" : "ЛОЖЬ") << "\n";
    }
    else if (c == "TCHECK" && pc == 3) {
        auto& t = trees[p[1]];
        if (p[2] == "complete") cout << (t.isComplete() ? "ИСТИНА" : "ЛОЖЬ") << "\n";
    }
    else if (c == "PRINT" && pc == 2) {
        string n = p[1];
        if      (arrays.count(n)) arrays[n].print();
        else if (slists.count(n)) slists[n].printForward();
        else if (dlists.count(n)) dlists[n].printForward();
        else if (stacks.count(n)) stacks[n].print();
        else if (queues.count(n)) queues[n].print();
        else if (trees.count(n))  trees[n].print();
        else cout << "не найдено\n";
    }
    else if (c == "DETAILS" && pc == 2) {
        string n = p[1];
        if      (arrays.count(n)) arrays[n].printDetails();
        else if (slists.count(n)) slists[n].printDetails();
        else if (dlists.count(n)) dlists[n].printDetails();
        else if (stacks.count(n)) stacks[n].printDetails();
        else if (queues.count(n)) queues[n].printDetails();
        else if (trees.count(n))  trees[n].printDetails();
        else cout << "не найдено\n";
    }
    else {
        cout << "Неизвестная команда: " << c << "\n";
        return 1;
    }

    if (changed) saveFile(filename);
    return 0;
}