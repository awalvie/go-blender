#!/bin/python3
import os
from markdownify import markdownify

TXT_DIR = 'txt/'
MD_DIR  = 'md/'
EXT_MD  = '.md'

class Markdown:
    def __init__(self, name, title, subtitle, created_at, html):
        self.name = name
        self.title = title
        self.subtitle = subtitle
        self.created_at = created_at[0:4] + '-' + created_at[4-6] + '-' + created_at[6:]
        self.markdown = None
        self.html = html

def parseFiles(files):
    markdownList = []

    for file in files:
        with open(TXT_DIR + file, 'r') as f:
            # extract data from the txt
            name = file.removesuffix('.txt')
            title = f.readline().strip()
            subtitle = f.readline().strip()
            created_at = f.readline().strip()
            html = f.read().strip()

            # created a markdown object from the data
            m = Markdown(name, title, subtitle, created_at, html)

            # append the markdown object to the list
            markdownList.append(m)

    # return object list
    return markdownList

def htmlToMarkdown(markdownList):
    for markdownFile in markdownList:
        markdownFile.markdown = markdownify(markdownFile.html)
    return markdownList

def renderMarkdown(markdownList):
    for markdownFile in markdownList:
        with open(MD_DIR + markdownFile.name + EXT_MD, 'w') as f:
            f.write('---\n')
            f.write('title: ' + markdownFile.title + '\n')
            f.write('subtitle: ' + markdownFile.subtitle + '\n')
            f.write('created: ' + markdownFile.created_at + '\n')
            f.write('---\n\n')
            f.write(markdownFile.markdown)

if __name__ == "__main__":
    # get list of files in the current directory
    files = os.listdir(TXT_DIR)

    # parse files and turn them into markdown objects
    markdownList = parseFiles(files)

    # translate the html and relevant frontmatter to compatible markdown
    htmlToMarkdown(markdownList)

    # write new markdown files
    renderMarkdown(markdownList)
