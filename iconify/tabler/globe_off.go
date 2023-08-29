package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlobeOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M7.353 7.355a4 4 0 0 0 5.29 5.293m2.007-2.009a4 4 0 0 0-5.3-5.284M5.75 15a8.015 8.015 0 0 0 9.792.557m2.02-1.998A8.015 8.015 0 0 0 15 2m-4 15v4m-4 0h8M3 3l18 18"/>`),
		g.Group(children),
	)
}