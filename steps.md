### Checklist of things to do

- [x] Info, Warn and Error Logging
- [x] Command line parsing
- [ ] Figure out markdown directory structure
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


- [ ] Initialize directory with a default structure

