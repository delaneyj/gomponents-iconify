package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PillCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopPillCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="m13.546 7.182l6.364 6.364a4.5 4.5 0 1 1-6.364 6.364l-6.364-6.364a4.5 4.5 0 1 1 6.364-6.364ZM16 12.5c-.667-.667-1.5-.833-2.5-.5c-1 .333-1.667 1.333-2 3l3.42 3.444l.006.006c.973.973 2.55.973 3.524 0l.007-.008a2.49 2.49 0 0 0-.015-3.52L16 12.5Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopPillCircleFilled0)"/></g>`),
		g.Group(children),
	)
}