package grommet_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FingerPrint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="2" d="M14 15a2 2 0 1 0-2 2h0m0 3a5 5 0 1 1 5-5a1.5 1.5 0 0 0 3 0a8 8 0 1 0-8 8h2M1 15c0 2.672.953 5.122 2.537 7.027M20.52 8.042A10.978 10.978 0 0 0 12 4a10.977 10.977 0 0 0-8.464 3.974m14.99-5.363A13.939 13.939 0 0 0 12 1a13.94 13.94 0 0 0-6.481 1.587"/>`),
		g.Group(children),
	)
}