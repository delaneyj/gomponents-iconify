package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CityThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M22 2h-7v5H9V2H2v20h20V2ZM9 9h6v11h-2v-4h-2v4H9V9ZM7 20H4V4h3v16Zm10 0V4h3v16h-3Zm-3-9h-4v2h4v-2Z"/>`),
		g.Group(children),
	)
}