### Checklist of things to do

- [x] Info, Warn and Error Logging
- [x] Command line parsing
- [x] Figure out markdown directory structure
	- Allow users to turn any page into a folder, ref: [xxiivv](https://wiki.xxiivv.com/site/home.html)

	```console
	|--
	index/
	|-- index.md
	|-- about/
		|-- index.md
		|-- friends.md
		|-- tools
			|-- index.md
			|-- have.md
			|-- want.md
	|-- knowledge/
		|-- index.md
		|-- knowledge.md
		|-- recollect.md
		|-- notions.md
		|-- quotes.md
		|-- recollect.md
	|-- exploration/
		|-- index.md
		|-- exploration.md
		|-- art.md
		|-- writing.md
		|-- language.md
	|-- compute/
		|-- index.md
		|-- languages.md
		|-- projects/
		|-- techlog/
		|-- devlog/
			|-- devlog 1
			|-- devlog 2
			|-- devlog 3
	...
	```

	```text
	1. Input path to folder
	2. List all files in folder
	3. Add each file/folder as child node
	4. If any folders in list --> Go back to step 2
	5. Render tree

	map: key=/index value=[compute, exploration, ...]
	map: key=/index/compute value=[_index, language, projects, techlog, devlog]
	map: key=/index/compute/devlog value=[log1, log2, log3]
	map: key=/index/compute/devlog/log1.md = []

	devlog.html

	|   devlog    | log       |

	// map filename -> path
	// path -> child

	[[ devlog ]]

	```
	-


- [x] Initialize directory with a default structure
- [x] Generate directory tree for source files
- [x] Translate lyceum html files to md, to serve as test subjects
- [x] Sort translated md files into folders
- [x] Parse source markdown files and render them in HTML using a library like [goldmark](https://github.com/yuin/goldmark)
- [x] Parse frontmatter
- [x] Now, how the hell am I going to do templating
	- Build steps are as follows:
		- Initialize `goldmark`
		- Iterate through all files and dirs
		- If file
			- Read file data
			- Convert to HTML
			- Parse metadata
			- Send HTML data, metadata and file map to template
			- Somehow render the navigation logic in the template
- [x] Create templates
- [x] Use frontmatter in templates
- [ ] Make it look good
- [ ] Extend markdown parsing to support backlinks
