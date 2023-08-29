package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Looks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 6A11 11 0 0 0 1 17h2c0-4.96 4.04-9 9-9s9 4.04 9 9h2A11 11 0 0 0 12 6m0 4c-3.86 0-7 3.14-7 7h2a5 5 0 0 1 5-5a5 5 0 0 1 5 5h2c0-3.86-3.14-7-7-7Z"/>`),
		g.Group(children),
	)
}