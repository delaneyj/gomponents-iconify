package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laundry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M16.9 13H16a2 2 0 0 1-2-2a2 2 0 1 1-4 0a2 2 0 0 1-2 2h-.9M5.5 4h3m.217 17h6.566a4 4 0 0 1 3.392 1.88l.075.12H21V1H3v22h2.25l.075-.12A4 4 0 0 1 8.717 21ZM12 17a5 5 0 1 1 0-10a5 5 0 0 1 0 10Z"/>`),
		g.Group(children),
	)
}