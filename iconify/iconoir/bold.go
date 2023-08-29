package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-width="1.5" d="M12 11.667H8m4 0s3.333 0 3.333-3.334C15.333 5 12 5 12 5H8.6a.6.6 0 0 0-.6.6v6.067m4 0s4 0 4 3.666C16 19 12 19 12 19H8.6a.6.6 0 0 1-.6-.6v-6.733"/>`),
		g.Group(children),
	)
}