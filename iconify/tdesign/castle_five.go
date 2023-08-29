package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CastleFive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M20 2v1h-4V2h-2v7.131l-2-1.333l-2 1.333V2H8v1H4V2H2v20h20V2h-2ZM4 5h4v5H6v10H4V5Zm4 15v-8h1.303L12 10.202L14.697 12H16v8h-3v-5h-2v5H8Zm10 0V10h-2V5h4v15h-2Z"/>`),
		g.Group(children),
	)
}