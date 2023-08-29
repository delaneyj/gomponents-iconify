package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FormatTextRotationDown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6 19.73l-3-3h2V4.27h2v12.46h2l-3 3m8-10.35v3.75l5.03-1.88L14 9.38M21 12l-11 4.73v-2.06l2.19-.94V8.77L10 7.83V5.77l11 4.73V12Z"/>`),
		g.Group(children),
	)
}