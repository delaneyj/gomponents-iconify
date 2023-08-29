package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MovedLocation(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m18 13l-1.4-1.4l2.15-2.15q-1.075-.2-3.438.225T10 12.975q.8-2.925 3.225-4.7T18.7 6.5l-2.1-2.1L18 3l5 5l-5 5Zm-6 9q-4.025-3.425-6.012-6.362T4 10.2q0-3.4 2.325-5.8T12 2q.875 0 1.75.188t1.675.562l-1.55 1.55q-.45-.15-.925-.225T12 4Q9.475 4 7.737 5.813T6 10.2q0 1.775 1.475 4.05T12 19.325q1.5-1.375 2.625-2.625t1.85-2.4l1.45 1.45q-1 1.475-2.475 3.025T12 22Z"/>`),
		g.Group(children),
	)
}