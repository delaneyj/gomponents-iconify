package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MoreVert(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M42.5 107q-17.5 0-30-12.5T0 64t12.5-30.5t30-12.5t30 12.5T85 64T72.5 94.5t-30 12.5zm0 42q17.5 0 30 12.5T85 192t-12.5 30.5t-30 12.5t-30-12.5T0 192t12.5-30.5t30-12.5zm0 128q17.5 0 30 12.5T85 320t-12.5 30.5t-30 12.5t-30-12.5T0 320t12.5-30.5t30-12.5z"/>`),
		g.Group(children),
	)
}