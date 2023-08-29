package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChevronBigDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.515 8.465L12 16.95l8.485-8.485L19.07 7.05L12 14.122L4.929 7.05L3.515 8.465Z"/>`),
		g.Group(children),
	)
}