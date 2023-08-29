package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ForbiddenCircleOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 1.25C6.063 1.25 1.25 6.063 1.25 12S6.063 22.75 12 22.75S22.75 17.937 22.75 12S17.937 1.25 12 1.25ZM2.75 12A9.25 9.25 0 0 1 12 2.75c2.284 0 4.376.828 5.99 2.2a.719.719 0 0 0-.02.02l-13 13a.912.912 0 0 0-.02.02A9.213 9.213 0 0 1 2.75 12Zm3.26 7.05A9.25 9.25 0 0 0 19.05 6.01a.733.733 0 0 1-.019.02l-13 13a.742.742 0 0 1-.02.018Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}