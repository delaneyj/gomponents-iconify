package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkMultipleTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M6.25 4a5.25 5.25 0 0 0-1.703 10.217c.06-.51.181-1 .353-1.467A3.751 3.751 0 0 1 6.25 5.5h6.5a3.75 3.75 0 1 1 0 7.5h-1a.75.75 0 0 0 0 1.5h1a5.25 5.25 0 1 0 0-10.5h-6.5Zm13.14 6.196c-.092.5-.243.98-.443 1.434A3.502 3.502 0 0 1 18 18.5h-7a3.5 3.5 0 1 1 0-7h1.25a.75.75 0 0 0 0-1.5H11a5 5 0 0 0 0 10h7a5 5 0 0 0 1.39-9.804Z"/>`),
		g.Group(children),
	)
}