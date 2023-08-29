package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngleDownCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopAngleDownCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="m19.64 11.268l-6 5c-1.024.854-2.304-.683-1.28-1.536l6-5c1.024-.854 2.305.683 1.28 1.536Z"/><path d="m7.64 9.732l6 5c1.024.853-.256 2.39-1.28 1.536l-6-5c-1.024-.853.256-2.39 1.28-1.536Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopAngleDownCircleFilled0)"/></g>`),
		g.Group(children),
	)
}