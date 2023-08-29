package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func O(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10 24c0-7.732 6.268-14 14-14s14 6.268 14 14s-6.268 14-14 14s-14-6.268-14-14Zm14-10c-5.523 0-10 4.477-10 10s4.477 10 10 10s10-4.477 10-10s-4.477-10-10-10Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}