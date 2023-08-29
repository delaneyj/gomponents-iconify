package bxs

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightTopArrowCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M19.071 4.929c-3.899-3.898-10.243-3.898-14.143 0c-3.898 3.899-3.898 10.244 0 14.143c3.899 3.898 10.243 3.898 14.143 0c3.899-3.9 3.899-10.244 0-14.143zm-3.536 10.607l-2.828-2.829l-3.535 3.536l-1.414-1.414l3.535-3.536l-2.828-2.829h7.07v7.072z"/>`),
		g.Group(children),
	)
}