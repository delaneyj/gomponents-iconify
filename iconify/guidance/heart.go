package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Heart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12 8.5V7.486m0 .264c0-2.9 2.35-5.25 5.25-5.25s5.25 2.382 5.25 5.282c0 1.56-.688 3.055-1.88 4.062l-5.657 4.776A8.35 8.35 0 0 0 12 23a8.35 8.35 0 0 0-2.963-6.38L3.38 11.844A5.327 5.327 0 0 1 1.5 7.782c0-2.9 2.35-5.282 5.25-5.282S12 4.85 12 7.75Z"/>`),
		g.Group(children),
	)
}