package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TwotoneDetails(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 8.92L18.6 19H13V8.92zm-2 0V19H5.4L11 8.92z" opacity=".3"/><path fill="currentColor" d="M12 3L2 21h20L12 3zm1 5.92L18.6 19H13V8.92zm-2 0V19H5.4L11 8.92z"/>`),
		g.Group(children),
	)
}