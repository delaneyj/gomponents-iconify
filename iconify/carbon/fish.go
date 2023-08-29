package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fish(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="24" cy="7" r="1" fill="currentColor"/><path fill="currentColor" d="M26 2a16.018 16.018 0 0 0-16 16v2H3a1 1 0 0 0-.707 1.707l8 8A1 1 0 0 0 12 29v-7h2A16.018 16.018 0 0 0 30 6V2ZM10 26.586L5.414 22H10ZM12 20v-2a13.94 13.94 0 0 1 3.908-9.678l7.77 7.77A13.94 13.94 0 0 1 14 20ZM28 6a13.916 13.916 0 0 1-2.98 8.605L17.395 6.98A13.916 13.916 0 0 1 26 4h2Z"/>`),
		g.Group(children),
	)
}