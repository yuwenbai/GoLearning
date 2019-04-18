#!/usr/bin/python
# -*- coding: UTF-8 -*-

import os
import sys
import plistlib


arg1 = sys.argv[1]   
print arg1
info = plistlib.readPlist("test.plist")
info['items'][0]['assets'][0]['url'] = arg1
print info['items'][0]['assets'][0]['url']



print "os.path.realpath(__file__)=%s" % os.path.realpath(__file__)

print "os.path.dirname(os.path.realpath(__file__))=%s" % os.path.dirname(os.path.realpath(__file__))

print "os.path.dirname(os.path.realpath(__file__))=%s" % os.path.dirname(os.path.realpath(__file__))+'/temp.plist'
# 创建临时文件
plistlib.writePlist(info,os.path.dirname(os.path.realpath(__file__))+'/temp.plist')
# 
curpath = os.path.dirname(os.path.realpath(__file__))
# 当前目录下如果有 manifest 文件 先干掉 cause os.rename会报错！！！垃圾
if os.path.isfile(os.path.join(curpath,'test.plist')):
	os.remove('test.plist')

os.rename('temp.plist', 'test.plist')

print '参数个数为:', len(sys.argv), '个参数。'
print '参数列表:', str(sys.argv)