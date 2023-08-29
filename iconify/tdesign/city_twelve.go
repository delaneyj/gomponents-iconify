package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CityTwelve(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 2H2v20h20V5h-8v5h-4V2ZM8 12h8v8h-3v-5h-2v5H8v-8Zm-2 8H4V4h4v6H6v10Zm12 0V10h-2V7h4v13h-2Z"/>`),
		g.Group(children),
	)
}