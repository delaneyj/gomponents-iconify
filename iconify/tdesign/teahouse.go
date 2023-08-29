package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Teahouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M8 2v1h8V2h2v1h1v2h-1v1.667L20.5 10H22v2h-1v8h1v2H2v-2h1v-8H2v-2h1.5L6 6.667V5H5V3h1V2h2Zm0 3v1h8V5H8Zm8.5 3h-9L6 10h12l-1.5-2Zm2.5 4h-3v8h3v-8Zm-5 8v-8h-4v8h4Zm-6 0v-8H5v8h3Z"/>`),
		g.Group(children),
	)
}