package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Folder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M1.5 10V2.5h5l3 3h11v3m3 .25V8.5H4.6l-.15.25l-.234.492A28 28 0 0 0 1.5 21.272v.228h19v-.128a28 28 0 0 1 2.757-12.116l.243-.506Z"/>`),
		g.Group(children),
	)
}