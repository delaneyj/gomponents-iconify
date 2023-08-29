package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JournalSharp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M290 32H104a24 24 0 0 0-24 24v400a24 24 0 0 0 24 24h186Zm118 0h-58v448h58a24 24 0 0 0 24-24V56a24 24 0 0 0-24-24Z"/>`),
		g.Group(children),
	)
}