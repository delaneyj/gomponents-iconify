package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M10 19v4.5m0-4.5c0-4 3.167-9.806 7.063-13.053L18.5 4.75V4.5h-4.973M10 19c0-4-3.167-9.806-7.063-13.053L1.5 4.75V4.5h12.027M10 23.5H5m5 0h5m-9.5-15h9m3.5 1a4.5 4.5 0 1 0-4.473-5"/>`),
		g.Group(children),
	)
}