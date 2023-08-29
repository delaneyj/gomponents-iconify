package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Teradata(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0C5 0 0 5.65 0 12.08C0 18.83 5 24 12 24s12-5.17 12-11.92C24 5.65 19 0 12 0M8.47 3.44h3.5V6.7h3.58v2.86h-3.58v5.22c0 1.58.77 2.27 1.93 2.27c.42 0 .98-.12 1.51-.32c.38 1 1.05 1.9 1.77 2.62a7 7 0 0 1-3.52.97c-3.12 0-5.19-1.65-5.19-5.28V3.45Z"/>`),
		g.Group(children),
	)
}