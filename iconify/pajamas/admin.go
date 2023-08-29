package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Admin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="m7.879 5l1.06-1.06l1.421-1.422a3.5 3.5 0 0 0-3.653 4.674l.326.897l-.675.674l-3.797 3.798a.621.621 0 1 0 .878.878l3.798-3.797l.674-.675l.897.325a3.5 3.5 0 0 0 4.674-3.653L12.06 7.062L11 8.12L9.94 7.06l-1-1L7.878 5Zm6.173-1.93A4.987 4.987 0 0 1 15 6a5 5 0 0 1-6.703 4.703L4.5 14.5a2.121 2.121 0 0 1-3-3l3.797-3.797A5 5 0 0 1 13 2l-1.076 1.076l-.863.863L10 5l1 1l1.06-1.06l.864-.864L14 3l.052.07Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}