package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LaunchStudyOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="23" cy="25" r="1" fill="currentColor"/><path d="M26 22v6H6v-6h20m0-2H6a2 2 0 0 0-2 2v6a2 2 0 0 0 2 2h20a2 2 0 0 0 2-2v-6a2 2 0 0 0-2-2z" fill="currentColor"/><path d="M20.59 10.59L17 14.17V2h-2v12.17l-3.59-3.58L10 12l6 6l6-6l-1.41-1.41z" fill="currentColor"/><path d="M8 24h12v2H8z" fill="currentColor"/>`),
		g.Group(children),
	)
}