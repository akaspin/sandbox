#!/usr/bin/env python
# wscript - Waf build script

top = 'src'
out = 'build'

def options(opt):
    pass

def configure(conf):
    conf.load('go', tooldir='tools')

def build(bld):
    bld.add_subdirs('src')