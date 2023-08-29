package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func R(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M71 128v64c42-40 99-64 161-64c19 0 38 3 55 7v75c-17-6-36-11-55-11c-87 0-159 70-161 157v298H0V128h71z"/>`),
		g.Group(children),
	)
}