# goscript
go扩展脚本支持,分别运行 Lua脚本 、 JavaScript脚本 、 Tengo脚本，比较运算map耗时

*扩展脚本：

--.js  JavaScript  AST  https://github.com/robertkrimen/otto
--.lua GopherLua   Lua5.1 VM  https://github.com/yuin/gopher-lua
--.tengo Tengo native GoVM      https://github.com/d5/tengo     Benchmark测试结果
--.py  go-python   C-API      https://github.com/sbinet/go-python  (未调用)

----
*测试结果：
----
Time consuming to create  50000 Maps :
        Lua Test :       9346 ms;
        JS Test :        9136 ms;
        Go Test :        42 ms;
        TenGo Test :     4263 ms;
        
执行速度: go>tengo>js>Lua>Python
