package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fullcube(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M15.5 3.03L4.7 9.263v12.47l10.8 6.237l10.8-6.234V9.266L15.5 3.028zm0 4l6.327 3.65l-6.327 3.654l-6.326-3.652L15.5 7.03zm9.488 3.57L16 15.79v10.377c0 .275-.225.5-.5.5s-.5-.225-.5-.5v-10.38l-8.987-5.19a.499.499 0 1 1 .5-.865l8.988 5.19l8.99-5.19a.501.501 0 0 1 .683.184a.502.502 0 0 1-.185.683z"/>`),
		g.Group(children),
	)
}