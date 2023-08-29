package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Export(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M24.086 20.904c-1.805 3.113-5.163 5.212-9.023 5.22A10.45 10.45 0 0 1 4.625 15.688A10.45 10.45 0 0 1 15.063 5.25c3.86.007 7.216 2.105 9.022 5.218l3.962 2.284l.143.082C26.88 6.784 21.504 2.25 15.063 2.248C7.64 2.25 1.625 8.265 1.623 15.688c.003 7.42 6.018 13.435 13.44 13.437c6.442-.002 11.82-4.538 13.127-10.59l-.14.082l-3.964 2.287zm4.314-5.216l-7.15-4.13v2.298H10.275v3.66H21.25v2.298l7.15-4.126z"/>`),
		g.Group(children),
	)
}