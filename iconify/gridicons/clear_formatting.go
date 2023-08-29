package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClearFormatting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m10.837 10.163l-4.6 4.6L10 4h4l.777 2.223l-2.144 2.144l-.627-2.092l-1.169 3.888zm5.495.506L19.244 19H15.82l-1.049-3.5H11.5L5 22l-1.5-1.5l17-17L22 5l-5.668 5.669zm-2.311 2.31l-.031.031l.032-.01l-.001-.021z"/>`),
		g.Group(children),
	)
}