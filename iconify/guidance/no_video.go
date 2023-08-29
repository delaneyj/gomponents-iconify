package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoVideo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M17.5 17.5v-3h.054c.94 0 1.877.274 2.564.917c.814.761 1.51 1.644 2.061 2.624L23 19.5h.5v-15H23l-.82 1.459a10.929 10.929 0 0 1-2.062 2.624c-.687.643-1.624.917-2.564.917H17.5v-5h-13m11.5 15H.5V4M.5.5l23 23"/>`),
		g.Group(children),
	)
}