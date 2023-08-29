package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListUnordered(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M9 17h10M9 12h10M9 7h10M5.002 17v.002H5V17h.002Zm0-5v.002H5V12h.002Zm0-5v.002H5V7h.002Z"/>`),
		g.Group(children),
	)
}