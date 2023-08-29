package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderOpen(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feFolderOpen0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feFolderOpen1" fill="currentColor" fill-rule="nonzero"><path id="feFolderOpen2" d="M11 5h8a2 2 0 0 1 2 2v12a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V5c0-1.1.9-2 2-2h4l2 2Zm-2 6v8h10v-8H9ZM5 7v12h2V9h12V7H5Z"/></g></g>`),
		g.Group(children),
	)
}