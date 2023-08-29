package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilePdf(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3 1h12.414L21 6.586V11h-2V9h-6V3H5v18h16v2H3V1Zm12 2.414V7h3.586L15 3.414ZM6 12h3.714c.71 0 1.286.576 1.286 1.286v2.428c0 .71-.576 1.286-1.286 1.286H8v3H6v-8Zm2 3h1v-1H8v1Zm3.5-3h3.714c.71 0 1.286.576 1.286 1.286v5.428c0 .71-.576 1.286-1.286 1.286H11.5v-8Zm2 2v4h1v-4h-1Zm3.5-.714c0-.71.576-1.286 1.286-1.286h3.38v2H19v1h2.667v2H19v3h-2v-6.714Z"/>`),
		g.Group(children),
	)
}