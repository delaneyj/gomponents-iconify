package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForkOutlined(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 2a2.993 2.993 0 0 0-1 5.816V11H6V7.816a3 3 0 1 0-2 0V11a2 2 0 0 0 2 2h5v4.184a3 3 0 1 0 2 0V13h5a2 2 0 0 0 2-2V7.816A2.993 2.993 0 0 0 19 2ZM5 6a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm7 15a1 1 0 1 1 1-1a1 1 0 0 1-1 1Zm7-15a1 1 0 1 1 1-1a1 1 0 0 1-1 1Z"/>`),
		g.Group(children),
	)
}