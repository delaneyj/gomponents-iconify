package wi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoonWaningGibbousOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 30 30"),
		g.Raw(`<path fill="currentColor" d="M3.74 14.49c0 1.22.19 2.4.56 3.54s.91 2.17 1.6 3.09s1.5 1.72 2.42 2.42s1.95 1.23 3.09 1.6s2.32.56 3.54.56c5.03-1.4 7.54-5.14 7.54-11.22c0-1.18-.14-2.3-.42-3.37s-.65-2.01-1.13-2.83s-1.04-1.57-1.68-2.24s-1.34-1.24-2.06-1.68s-1.47-.81-2.26-1.07c-1.52 0-2.98.3-4.37.89S8.02 5.57 7.02 6.57s-1.8 2.19-2.39 3.57s-.89 2.83-.89 4.35z"/>`),
		g.Group(children),
	)
}