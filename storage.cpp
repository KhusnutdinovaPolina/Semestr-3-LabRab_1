#include "storage.h"

map<string, Array> arrays;
map<string, SList> slists;
map<string, DList> dlists;
map<string, Stack> stacks;
map<string, Queue> queues;
map<string, CBT>   trees;

void saveFile(const string& filename) {
    ofstream f(filename);
    for (auto& kv : arrays) { f << "M " << kv.first << ":"; kv.second.save(f); f << "\n"; }
    for (auto& kv : slists) { f << "F " << kv.first << ":"; kv.second.save(f); f << "\n"; }
    for (auto& kv : dlists) { f << "L " << kv.first << ":"; kv.second.save(f); f << "\n"; }
    for (auto& kv : stacks) { f << "S " << kv.first << ":"; kv.second.save(f); f << "\n"; }
    for (auto& kv : queues) { f << "Q " << kv.first << ":"; kv.second.save(f); f << "\n"; }
    for (auto& kv : trees)  { f << "T " << kv.first << ":"; kv.second.save(f); f << "\n"; }
}

void loadFile(const string& filename) {
    ifstream f(filename);
    if (!f.is_open()) return;
    string line;
    while (getline(f, line)) {
        size_t c = line.find(':');
        if (c == string::npos) continue;

        string head[100]; int hc = 0;
        splitWords(line.substr(0, c), head, hc);
        if (hc < 2) continue;

        string type = head[0];
        string name = head[1];

        istringstream vals(line.substr(c + 1));

        if      (type == "M") arrays[name].load(vals);
        else if (type == "F") slists[name].load(vals);
        else if (type == "L") dlists[name].load(vals);
        else if (type == "S") stacks[name].load(vals);
        else if (type == "Q") queues[name].load(vals);
        else if (type == "T") trees[name].load(vals);
    }
}