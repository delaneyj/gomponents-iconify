package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VideoCameraOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 4h17v4.434l7-4.2v15.49l-7-4V20H0V4Zm17 9.42l5 2.857v-8.51l-5 3v2.653ZM2 6v12h13V6H2Zm2 2h6v2H4V8Z"/>`),
		g.Group(children),
	)
}