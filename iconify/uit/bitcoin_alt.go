package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BitcoinAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M15.722 11.548A3.735 3.735 0 0 0 13.25 5H12V3.5a.5.5 0 0 0-1 0V5H8V3.5a.5.5 0 0 0-1 0V5H5.5a.5.5 0 0 0 0 1H7v12H5.5a.5.5 0 0 0 0 1H7v1.5a.5.5 0 0 0 1 0V19h3v1.5a.5.5 0 0 0 1 0V19h3.25a3.74 3.74 0 0 0 .472-7.452zM8 6h5.25a2.75 2.75 0 1 1 0 5.5H8V6zm7.25 12H8v-5.5h7.25a2.75 2.75 0 1 1 0 5.5z"/>`),
		g.Group(children),
	)
}