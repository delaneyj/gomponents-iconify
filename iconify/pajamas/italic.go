package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Italic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.5 2a.75.75 0 0 0 0 1.5h1.93l-2.412 9H4A.75.75 0 0 0 4 14h5.5a.75.75 0 0 0 0-1.5H7.57l2.412-9H12A.75.75 0 0 0 12 2H6.5Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}