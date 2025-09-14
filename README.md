# kvstore - Redis Clone (in Go)

This is a small side project where I try to build a very simple version of Redis using Go.  
I mainly did this just to understand how Redis kinda works behind the scenes.  

Right now it can do really basic stuff like:

- store a value for a key
- get back the value for a key
- maybe a couple of extra toy commands

That’s about it. Don’t expect much, this is not production ready (please don’t actually use it in production lol).

Aiming to use and understand go core concepts like mutex, interface, embedded structs, TCP client-server implementation, file I/O operations for persistency, queues,concurrency, context, etc.  

## How to run

Clone this repo:

```bash
git clone https://github.com/dheeraj-mishra/kvstore.git
cd kvstore
make build
