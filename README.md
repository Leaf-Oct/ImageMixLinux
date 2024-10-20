# ImageMix for Linux

本人项目[ImageMix](https://github.com/Leaf-Oct/ImageMix)的Linux版，主要用于Ubuntu，其它发行版和桌面环境未测试过，不确定可用性

图片像素混淆器，在[米缸师兄的项目https://github.com/GensokyoOtaku/ImageMix](https://github.com/GensokyoOtaku/ImageMix)的基础上，借用了其思想和实现思路，用go语言实现了一个类似的程序，主要功能如下：

- 读取剪切板内的图片，随机修改10个像素信息为随机值，再以png格式写回剪切板内
- 保存剪切板内的图像(无论是修没修改过的)

该项目仅能编译和运行于Linux系统，因为暂时不熟悉go语言的跨平台开发，因此没有将Windows系统对应的代码写在一起。