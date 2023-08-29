package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dropbox(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m3 6.2l5 3.19l5-3.19L8 3L3 6.2m10 0l5 3.19l5-3.19L18 3l-5 3.2M3 12.55l5 3.19l5-3.19l-5-3.2l-5 3.2m15-3.2l-5 3.2l5 3.19l5-3.19l-5-3.2M8.03 16.8l5.01 3.2l5-3.2l-5-3.19l-5.01 3.19Z"/>`),
		g.Group(children),
	)
}