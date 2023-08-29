package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HandPointingLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M188 90a25.87 25.87 0 0 0-14.59 4.49A26 26 0 0 0 134 78.1V44a26 26 0 0 0-52 0v87l-7.53-12.1a26 26 0 0 0-45 26.07l4.67 8.25c34 60 48.07 84.77 93.86 84.77a86.1 86.1 0 0 0 86-86V116a26 26 0 0 0-26-26Zm14 62a74.09 74.09 0 0 1-74 74c-38.8 0-50-19.83-83.42-78.69L39.89 139A14 14 0 0 1 45 119.88a13.87 13.87 0 0 1 7-1.88a14 14 0 0 1 12.15 7l.1.17l18.68 30A6 6 0 0 0 94 152V44a14 14 0 0 1 28 0v68a6 6 0 0 0 12 0v-12a14 14 0 0 1 28 0v20a6 6 0 0 0 12 0v-4a14 14 0 0 1 28 0Z"/>`),
		g.Group(children),
	)
}