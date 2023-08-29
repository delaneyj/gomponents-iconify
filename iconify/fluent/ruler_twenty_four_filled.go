package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 4.25A2.25 2.25 0 0 1 9.25 2h5.5A2.25 2.25 0 0 1 17 4.25v15.5A2.25 2.25 0 0 1 14.75 22h-5.5A2.25 2.25 0 0 1 7 19.75V4.25ZM8.5 5v1.5h1.75a.75.75 0 0 0 0-1.5H8.5Zm0 3v1.5h3.75a.75.75 0 0 0 0-1.5H8.5Zm0 3.25v1.5h1.75a.75.75 0 0 0 0-1.5H8.5Zm0 3.25V16h3.75a.75.75 0 0 0 0-1.5H8.5Zm0 3V19h1.75a.75.75 0 0 0 0-1.5H8.5Z"/>`),
		g.Group(children),
	)
}