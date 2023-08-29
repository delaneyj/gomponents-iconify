package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiStrengthOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3C7.79 3 3.7 4.41.38 7C4.41 12.06 7.89 16.37 12 21.5c4.08-5.08 8.24-10.26 11.65-14.5C20.32 4.41 16.22 3 12 3m0 2c3.07 0 6.09.86 8.71 2.45l-5.1 6.36A8.432 8.432 0 0 0 12 13c-1.25 0-2.5.28-3.61.8L3.27 7.44C5.91 5.85 8.93 5 12 5Z"/>`),
		g.Group(children),
	)
}