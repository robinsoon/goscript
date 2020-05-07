# goscript
go扩展脚本支持,分别运行 Lua脚本 、 JavaScript脚本 、 Tengo脚本 与 go原生程序比较运算map耗时

*   **扩展脚本**

-  [.js]  JavaScript  AST  https://github.com/robertkrimen/otto  <br/>
-  [.lua] GopherLua   Lua5.1 VM  https://github.com/yuin/gopher-lua <br/>
-  [.tengo] Tengo native GoVM      https://github.com/d5/tengo     Benchmark测试结果<br/>
-  [.py]  go-python   C-API      https://github.com/sbinet/go-python  (未调用)<br/>

----
*   **测试结果**
----
Time consuming to create  50000 Maps :<br/>
        Lua Test :       9346 ms;<br/>
        JS Test :        9136 ms;<br/>
        Go Test :        42 ms;<br/>
        TenGo Test :     4263 ms;<br/>
<br/><br/>        
执行速度: go>tengo>js>Lua>Python<br/>
