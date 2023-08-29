package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Term(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 26h-3v-2h3V8h-3V6h3a2.002 2.002 0 0 1 2 2v16a2.003 2.003 0 0 1-2 2Z"/><circle cx="23" cy="16" r="2" fill="currentColor"/><circle cx="16" cy="16" r="2" fill="currentColor"/><circle cx="9" cy="16" r="2" fill="currentColor"/><path fill="currentColor" d="M7 26H4a2.002 2.002 0 0 1-2-2V8a2.002 2.002 0 0 1 2-2h3v2H4v16h3Z"/>`),
		g.Group(children),
	)
}