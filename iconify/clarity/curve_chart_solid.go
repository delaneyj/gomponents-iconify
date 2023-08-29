package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CurveChartSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M34 7v22a2 2 0 0 1-2 2H4a2 2 0 0 1-2-2V7a2 2 0 0 1 2-2h28a2 2 0 0 1 2 2Zm-21 5c1.817 0 2.674 1.499 4.039 6.275c.774 2.709 1.162 3.843 1.843 5.034C19.87 25.038 21.205 26 23 26h6a1 1 0 0 0 0-2h-6c-1.817 0-2.674-1.499-4.038-6.275c-.774-2.709-1.163-3.843-1.844-5.034C16.13 10.962 14.795 10 13 10H7a1 1 0 0 0 0 2Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}