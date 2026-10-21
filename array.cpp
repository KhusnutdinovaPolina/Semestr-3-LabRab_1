#include "array.h"

Array::Array() {
    capacity = 10;
    data = new string[capacity];
    size = 0;
}

Array::~Array() {
    delete[] data;
}

void Array::pushBack(const string& v) {
    if (size >= capacity) {
        capacity *= 2;
        string* bigger = new string[capacity];
        for (int i = 0; i < size; i++) bigger[i] = data[i];
        delete[] data;
        data = bigger;
    }
    data[size++] = v;
}

void Array::insertAt(int index, const string& v) {
    if (index < 0 || index > size) return;
    if (size >= capacity) {
        capacity *= 2;
        string* bigger = new string[capacity];
        for (int i = 0; i < size; i++) bigger[i] = data[i];
        delete[] data;
        data = bigger;
    }
    for (int i = size; i > index; i--) data[i] = data[i - 1];
    data[index] = v;
    size++;
}

string Array::get(int index) {
    if (index < 0 || index >= size) return "";
    return data[index];
}

void Array::set(int index, const string& v) {
    if (index >= 0 && index < size) data[index] = v;
}

void Array::removeAt(int index) {
    if (index < 0 || index >= size) return;
    for (int i = index; i < size - 1; i++) data[i] = data[i + 1];
    size--;
}

int Array::length() {
    return size;
}

void Array::clear() {
    size = 0;
}

void Array::print() {
    cout << "Массив:  [";
    for (int i = 0; i < size; i++) {
        cout << data[i];
        if (i < size - 1) cout << ", ";
    }
    cout << "]\n";
}

void Array::printDetails() {
    cout << "Массив (адреса ячеек):\n";
    for (int i = 0; i < size; i++) {
        cout << "  [" << i << "] addr=" << &data[i] << " value=\"" << data[i] << "\"\n";
    }
}

void Array::save(ostream& os) {
    for (int i = 0; i < size; i++) os << " " << data[i];
}

void Array::load(istream& is) {
    string v;
    while (is >> v) pushBack(v);
}