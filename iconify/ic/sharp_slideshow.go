package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpSlideshow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M10 8v8l5-4l-5-4zm11-5H3v18h18V3zm-2 16H5V5h14v14z"/>`),
		g.Group(children),
	)
}