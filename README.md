# Arrows

It's a game that has an idea similar to Redstone in Minecraft, idea stolen from [Onigiri](https://github.com/ArtemOnigiri/).

It has basic logic gates like AND, NOT, XOR, etc.

Work in progress...

## Technical details

- CORE - is main module/library, it contains logic - you can use it with Go. It almost done, but still have (a lot of) issues.
- WASM - Port for Web/JS. Kinda done, but isn't really tested
- SHARED - Shared library(dll/so). You could use it in C, C++, Python, etc. Not tested at all, and it has low priority, so...

</br>
In future, we'll try to make this game multithreading and maybe even decentralized (a map will be calculated on several computers).
</br>
</br>
I hope I'll make docs about adding new types of cells, compiling it, porting, etc
