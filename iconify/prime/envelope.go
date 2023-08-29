package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Envelope(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19 4.25H5A1.76 1.76 0 0 0 3.25 6v12A1.76 1.76 0 0 0 5 19.75h14A1.76 1.76 0 0 0 20.75 18V6A1.76 1.76 0 0 0 19 4.25ZM5 5.75h14a.25.25 0 0 1 .25.25v1.54L12 11.16L4.75 7.54V6A.25.25 0 0 1 5 5.75Zm14 12.5H5a.25.25 0 0 1-.25-.25V9.21l6.91 3.46a.76.76 0 0 0 .68 0l6.91-3.46V18a.25.25 0 0 1-.25.25Z"/>`),
		g.Group(children),
	)
}