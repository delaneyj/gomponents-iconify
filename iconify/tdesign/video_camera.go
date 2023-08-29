package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCamera(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 4h17v4.434l7-4.2v15.532l-7-4.2V20H0V4Zm15 2H2v12h13V6Zm2 7.234l5 3V7.766l-5 3v2.468Z"/>`),
		g.Group(children),
	)
}