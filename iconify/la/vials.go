package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Vials(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M5 5v2h1v12.5C6 21.4 7.6 23 9.5 23s3.5-1.6 3.5-3.5V7h1V5H5zm13 0v2h1v12.5c0 1.9 1.6 3.5 3.5 3.5s3.5-1.6 3.5-3.5V7h1V5h-9zM8 7h3v6H8V7zm13 0h3v6h-3V7zM8 15h3v4.5c0 .8-.7 1.5-1.5 1.5S8 20.3 8 19.5V15zm13 0h3v4.5c0 .8-.7 1.5-1.5 1.5s-1.5-.7-1.5-1.5V15zM2 25v2h28v-2H2z"/>`),
		g.Group(children),
	)
}