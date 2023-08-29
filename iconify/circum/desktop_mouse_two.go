package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DesktopMouseTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.437 21.938h-2.874a6.508 6.508 0 0 1-6.5-6.5V8.562a6.508 6.508 0 0 1 6.5-6.5h2.874a6.508 6.508 0 0 1 6.5 6.5v6.876a6.508 6.508 0 0 1-6.5 6.5ZM10.563 3.062a5.506 5.506 0 0 0-5.5 5.5v6.876a5.506 5.506 0 0 0 5.5 5.5h2.874a5.506 5.506 0 0 0 5.5-5.5V8.562a5.506 5.506 0 0 0-5.5-5.5Z"/><path fill="currentColor" d="M11.5 6.541v4a.5.5 0 0 0 1 0v-4a.5.5 0 0 0-1 0Z"/>`),
		g.Group(children),
	)
}