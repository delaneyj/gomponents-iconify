package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FigmaBroken(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 8.667H8.667m3.333 0V2H8.667a3.333 3.333 0 0 0 0 6.667m3.333 0v6.666H8.667M12 8.667h3.333m-6.666 0a3.333 3.333 0 0 0 0 6.666m0 0A3.333 3.333 0 1 0 12 18.667m3.333-10a3.333 3.333 0 1 0 0 6.666a3.333 3.333 0 0 0 0-6.666Zm0 0a3.333 3.333 0 0 0 0-6.667"/>`),
		g.Group(children),
	)
}