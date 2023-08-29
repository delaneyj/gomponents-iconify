package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7.5 13h5a.5.5 0 0 1 .09.992L12.5 14h-5a.5.5 0 0 1-.09-.992L7.5 13h5h-5Zm-2-4h9a.5.5 0 0 1 .09.992L14.5 10h-9a.5.5 0 0 1-.09-.992L5.5 9h9h-9Zm-2-4h13a.5.5 0 0 1 .09.992L16.5 6h-13a.5.5 0 0 1-.09-.992L3.5 5h13h-13Z"/>`),
		g.Group(children),
	)
}