package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TextEffectsGaTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M17 2.5A2.5 2.5 0 0 1 19.5 5v2.5a2.5 2.5 0 0 1 0 5V18a2.5 2.5 0 0 1-5 0V5A2.5 2.5 0 0 1 17 2.5zM18 18v-7h1.5a1 1 0 1 0 0-2H18V5a1 1 0 1 0-2 0v13a1 1 0 1 0 2 0zM2.5 8A2.5 2.5 0 0 1 5 5.5h6A2.5 2.5 0 0 1 13.5 8c0 1.837-.333 4.314-1.521 6.426C10.709 16.685 8.44 18.5 5 18.5a2.5 2.5 0 0 1 0-5c.74 0 1.25-.174 1.624-.411c.383-.243.712-.609.997-1.115c.243-.431.433-.934.573-1.474H5A2.5 2.5 0 0 1 2.5 8zm7.236 2.5c-.17.779-.432 1.542-.808 2.21C8.177 14.045 6.984 15 5 15a1 1 0 1 0 0 2c2.816 0 4.623-1.445 5.672-3.31C11.687 11.886 12 9.695 12 8a1 1 0 0 0-1-1H5a1 1 0 0 0 0 2h4.959c-.041.493-.113 1-.223 1.5z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}