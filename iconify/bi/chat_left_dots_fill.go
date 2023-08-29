package bi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatLeftDotsFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M0 2a2 2 0 0 1 2-2h12a2 2 0 0 1 2 2v8a2 2 0 0 1-2 2H4.414a1 1 0 0 0-.707.293L.854 15.146A.5.5 0 0 1 0 14.793V2zm5 4a1 1 0 1 0-2 0a1 1 0 0 0 2 0zm4 0a1 1 0 1 0-2 0a1 1 0 0 0 2 0zm3 1a1 1 0 1 0 0-2a1 1 0 0 0 0 2z"/>`),
		g.Group(children),
	)
}