package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Shovel(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M994 994q-43 37-113.5 29T738 981t-113-74l-94-94q-19-20-19-47.5t19-46.5l48-48l-387-388l-82 82q-19 19-45.5 19T19 365T0 319.5T19 274L274 19q19-19 45.5-19T365 19t19 45.5t-19 45.5l-82 82l388 387l48-48q19-19 46.5-19t47.5 19l94 94q40 41 74 113t42 142.5T994 994z"/>`),
		g.Group(children),
	)
}