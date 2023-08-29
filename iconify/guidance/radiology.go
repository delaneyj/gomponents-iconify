package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radiology(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M7.5 16s2-1 4.5-1m0 0c2.5 0 4.5 1 4.5 1M12 15v7m-4.5-1s2-1 4.5-1s4.5 1 4.5 1m-10-2.5s2.5-1 5.5-1s5.5 1 5.5 1m1-7.5v-1S16 8.5 12 8.5S5.5 10 5.5 10v1m6.286-4.498s-2.29-1.5-2.29-3.374c0-1.45 1.121-2.625 2.505-2.625c1.383 0 2.499 1.175 2.499 2.625c0 1.875-2.284 3.374-2.284 3.374h-.43ZM21.5 12.5h-19v.25S3 15 3 18s-.5 5.25-.5 5.25v.25h19v-.25S21 21 21 18s.5-5.25.5-5.25v-.25Z"/>`),
		g.Group(children),
	)
}