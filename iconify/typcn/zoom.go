package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zoom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 4c-3.859 0-7 3.141-7 7c0 .763.127 1.495.354 2.183l-.749.75l-.511.512l-1.008 1.045a3.076 3.076 0 0 0-.891 2.185a3.134 3.134 0 0 0 3.13 3.131c.757 0 1.504-.278 2.104-.784l.064-.055l.061-.061l1.512-1.51l.75-.749A6.983 6.983 0 0 0 13 18c3.859 0 7-3.141 7-7s-3.141-7-7-7zm0 12c-2.757 0-5-2.243-5-5s2.243-5 5-5s5 2.243 5 5s-2.243 5-5 5zm0-9c-2.205 0-4 1.794-4 4s1.795 4 4 4s4-1.794 4-4s-1.795-4-4-4zm0 7a3.001 3.001 0 0 1 0-6a3.001 3.001 0 0 1 0 6z"/>`),
		g.Group(children),
	)
}