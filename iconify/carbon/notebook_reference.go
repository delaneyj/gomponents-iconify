package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NotebookReference(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M4 20v2h3.586L2 27.586L3.414 29L9 23.414V27h2v-7H4zm15-10h7v2h-7zm0 5h7v2h-7zm0 5h7v2h-7z"/><path fill="currentColor" d="M28 5H4a2.002 2.002 0 0 0-2 2v10h2V7h11v20h13a2.002 2.002 0 0 0 2-2V7a2.002 2.002 0 0 0-2-2ZM17 25V7h11l.002 18Z"/>`),
		g.Group(children),
	)
}