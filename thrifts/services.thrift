include 'common.thrift'

namespace go thrift_demo

#use : thrift -out .. -r --gen go services.thrift to generate interface

service Greeter{
    common.Status greet(1:string name),
}