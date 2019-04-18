#!/usr/bin/python
# -*- coding: UTF-8 -*-

import os
import sys
import plistlib


print "os.path.realpath1 (__file__)=%s" % os.path.realpath(__file__)

arg1 = sys.argv[1]   
print "param from binary =%s" % arg1


curpath = os.path.dirname(os.path.realpath(__file__))
print curpath


info = plistlib.readPlist(curpath + "/manifest.plist")
info['items'][0]['assets'][0]['url'] = "https://update.microcontract.io/static/versions/bo-ios/" + arg1
print info['items'][0]['assets'][0]['url']



print "os.path.realpath(__file__)=%s" % os.path.realpath(__file__)

print "os.path.dirname(os.path.realpath(__file__))=%s" % os.path.dirname(os.path.realpath(__file__))

print "os.path.dirname(os.path.realpath(__file__))=%s" % os.path.dirname(os.path.realpath(__file__))+'/temp.plist'
# 创建临时文件
plistlib.writePlist(info,os.path.dirname(os.path.realpath(__file__))+'/temp.plist')
# 
curpath = os.path.dirname(os.path.realpath(__file__))
# 当前目录下如果有 manifest 文件 先干掉 cause os.rename会报错！！！垃圾
if os.path.isfile(os.path.join(curpath,'manifest.plist')):
	os.remove(curpath + "/manifest.plist")
print "目录为: %s"%os.listdir(os.getcwd())
os.rename(curpath + '/temp.plist', curpath + '/manifest.plist')

print '参数个数为:', len(sys.argv), '个参数。'
print '参数列表:', str(sys.argv)