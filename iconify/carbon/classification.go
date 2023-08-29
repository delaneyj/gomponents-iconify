package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Classification(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="15" cy="19" r="1" fill="currentColor"/><path fill="currentColor" d="m27.7 9.3l-7-7A.909.909 0 0 0 20 2H10a2.006 2.006 0 0 0-2 2v10H6a2.002 2.002 0 0 0-2 2v6a2.002 2.002 0 0 0 2 2h2v4a2.006 2.006 0 0 0 2 2h16a2.006 2.006 0 0 0 2-2V10a.91.91 0 0 0-.3-.7ZM20 4.4l5.6 5.6H20ZM6 16h9.597L19 19l-3.393 3H6Zm20 12H10v-4h5.609a2.008 2.008 0 0 0 1.313-.493l3.384-2.991a2.002 2.002 0 0 0 .01-3.023l-3.404-3.009A1.996 1.996 0 0 0 15.61 14H10V4h8v6a2.006 2.006 0 0 0 2 2h6Z"/>`),
		g.Group(children),
	)
}