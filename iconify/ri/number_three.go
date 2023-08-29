package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 2v1.362L12.81 9.55a6.501 6.501 0 1 1-7.117 8.028l1.941-.486A4.502 4.502 0 0 0 16.5 16a4.5 4.5 0 0 0-6.505-4.03l-.228.122l-.69-1.207L14.856 4H6.502V2H18Z"/>`),
		g.Group(children),
	)
}