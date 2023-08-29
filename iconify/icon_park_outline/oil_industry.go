package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OilIndustry(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linejoin="round" stroke-width="4"><path d="M10 6h28v36H10z"/><path stroke-linecap="round" d="M20.643 23.889c1.431-1.88 2.535-4.479 3.131-5.889c1.044 1.41 3.31 4.948 4.026 6.829c.894 2.35-1.342 5.171-4.026 5.171c-2.684 0-4.92-3.76-3.131-6.111ZM6 6h36M6 42h36M6 24h4m28 0h4"/></g>`),
		g.Group(children),
	)
}