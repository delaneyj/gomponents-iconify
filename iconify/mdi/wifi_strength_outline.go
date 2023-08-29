package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiStrengthOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 3C7.79 3 3.7 4.41.38 7H.36C4.24 11.83 8.13 16.66 12 21.5c3.89-4.84 7.77-9.67 11.64-14.5h.01C20.32 4.41 16.22 3 12 3m0 2c3.07 0 6.09.86 8.71 2.45L12 18.3L3.27 7.44C5.9 5.85 8.92 5 12 5Z"/>`),
		g.Group(children),
	)
}