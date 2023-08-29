package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Actions(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="2" d="M12 17.5a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11Zm0-11V1m0 22v-5.5M1 12h5.5m11 0H23M4.437 4.437l4.125 4.125m6.876 6.876l4.124 4.124m0-15.125l-4.125 4.125m-6.874 6.876l-4.126 4.124"/>`),
		g.Group(children),
	)
}