package majesticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ViewColumns(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M7 5H5a3 3 0 0 0-3 3v8a3 3 0 0 0 3 3h2V5zm2 14h6V5H9v14zm8-14v14h2a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3h-2z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}