package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hourglass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5" d="M12 12a7 7 0 0 0 7-7H5a7 7 0 0 0 7 7Zm0 0a7 7 0 0 1 7 7H5a7 7 0 0 1 7-7ZM5 2h14M5 22h14"/>`),
		g.Group(children),
	)
}