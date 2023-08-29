package wi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonWaningGibbousFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 30 30"),
		g.Raw(`<path fill="currentColor" d="M3.74 14.47c0 2.03.5 3.91 1.51 5.63s2.37 3.09 4.09 4.09s3.6 1.51 5.63 1.51c2.17-2.75 3.25-6.5 3.25-11.24c0-3.96-1.08-7.71-3.25-11.24c-2.03 0-3.91.5-5.63 1.5S6.26 7.1 5.25 8.83s-1.51 3.61-1.51 5.64z"/>`),
		g.Group(children),
	)
}