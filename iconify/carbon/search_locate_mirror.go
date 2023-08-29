package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SearchLocateMirror(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m21 28.6l-4.7-4.7c1.1-1.4 1.7-3.1 1.7-4.9c0-4.4-3.6-8-8-8s-8 3.6-8 8s3.6 8 8 8c1.8 0 3.5-.6 4.9-1.7l4.7 4.7l1.4-1.4zM10 25c-3.3 0-6-2.7-6-6s2.7-6 6-6s6 2.7 6 6s-2.7 6-6 6zm12-13h8v2h-8zm-8-5h16v2H14zm0-5h16v2H14z"/>`),
		g.Group(children),
	)
}