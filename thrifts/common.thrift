namespace go thrift_demo

struct Status {
    1: i8 code
    2: optional map<string, string> extra
}