package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlagOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M5.5 20.5a1 1 0 0 1-1-1v-15a1 1 0 0 1 1-1h6.38a1 1 0 0 1 .9.55l.72 1.45h5a1 1 0 0 1 1 1v8a1 1 0 0 1-1 1h-5.39a1 1 0 0 1-.89-.55l-.72-1.45h-5v6a1 1 0 0 1-1 1Zm1-15v6h6l1 2h4v-6h-5l-1-2h-5Z"/>`),
		g.Group(children),
	)
}