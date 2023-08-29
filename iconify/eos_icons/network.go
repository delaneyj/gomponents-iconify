package eos_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Network(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 4a8 8 0 1 0 8 8a8.003 8.003 0 0 0-8-8Zm-6 8a5.997 5.997 0 0 1 .105-1.095L9.6 14.4v.8a1.605 1.605 0 0 0 1.6 1.6v1.14A6.004 6.004 0 0 1 6 12Zm9.2 3.2h-.8v-2.4a.802.802 0 0 0-.8-.8H8.8v-1.6h1.6a.802.802 0 0 0 .8-.8V8h1.6a1.6 1.6 0 0 0 1.59-1.5a5.985 5.985 0 0 1 2.137 9.426A1.57 1.57 0 0 0 15.2 15.2ZM13 1a1 1 0 1 1-1-1a1 1 0 0 1 1 1Zm5 1a1 1 0 1 1-1-1a1 1 0 0 1 1 1Zm4 4a1 1 0 1 1-1-1a1 1 0 0 1 1 1Zm2 6a1 1 0 1 1-1-1a1 1 0 0 1 1 1Zm-2 6a1 1 0 1 1-1-1a1 1 0 0 1 1 1Zm-4 4a1 1 0 1 1-1-1a1 1 0 0 1 1 1Zm-5 1a1 1 0 1 1-1-1a1 1 0 0 1 1 1Zm-5-1a1 1 0 1 1-1-1a1 1 0 0 1 1 1Zm-4-4a1 1 0 1 1-1-1a1 1 0 0 1 1 1Zm-2-6a1 1 0 1 1-1-1a1 1 0 0 1 1 1Zm2-6a1 1 0 1 1-1-1a1 1 0 0 1 1 1Zm4-4a1 1 0 1 1-1-1a1 1 0 0 1 1 1Z"/>`),
		g.Group(children),
	)
}