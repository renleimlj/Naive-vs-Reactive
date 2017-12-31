# Naive-vs-Reactive

## Reactive 动机

* 对刺激反应灵敏

    * 事件驱动

    * 对加载作出反应，可拓展

    * 对失败作出反应，有弹性

    * 对用户作出反应，响应式

## 问题描述

![](https://jersey.github.io/documentation/latest/images/rx-client-problem.png)

## 同步方案

![](https://jersey.github.io/documentation/latest/images/rx-client-sync-approach.png)

## 异步方案

![](https://jersey.github.io/documentation/latest/images/rx-client-async-approach.png)

## 运行截图

![](https://s1.ax1x.com/2017/12/31/pSsDtH.png)

如上图所示，使用异步比单纯地使用同步要快很多，异步只需约745ms，而同步则需约5.57s。