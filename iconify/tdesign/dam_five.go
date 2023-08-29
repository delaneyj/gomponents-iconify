package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DamFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2 2h20v2h-1v16h1v2H2v-2h1V4H2V2Zm3 2v16h2V4H5Zm4 0v7c.836-.628 1.874-1 3-1s2.164.372 3 1V4H9Zm8 0v16h2V4h-2Zm-2 16v-5a3 3 0 1 0-6 0v5h6Z"/>`),
		g.Group(children),
	)
}