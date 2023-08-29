package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperPlaneTiltLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M225.88 30.12a13.83 13.83 0 0 0-13.7-3.58h-.11L20.14 84.77A14 14 0 0 0 18 110.88l85.61 40.55a2.08 2.08 0 0 1 .95.95L145.12 238a13.87 13.87 0 0 0 12.61 8c.4 0 .81 0 1.21-.05a13.9 13.9 0 0 0 12.29-10.09l58.2-191.93v-.11a13.83 13.83 0 0 0-3.55-13.7Zm-8 10.4l-58.15 191.91v.11a1.88 1.88 0 0 1-1.76 1.45a1.86 1.86 0 0 1-2-1.14l-40-84.36l48.24-48.24a6 6 0 1 0-8.49-8.49L107.52 140l-84.37-40a2 2 0 0 1 .31-3.74h.11l191.91-58.18a1.94 1.94 0 0 1 1.92.52a2 2 0 0 1 .52 1.92Z"/>`),
		g.Group(children),
	)
}