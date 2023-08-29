package bx

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BxsRightDownArrowCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path d="M4.929 4.929c-3.898 3.899-3.898 10.244 0 14.143c3.899 3.898 10.243 3.898 14.143 0c3.898-3.899 3.898-10.244 0-14.143c-3.9-3.899-10.244-3.899-14.143 0zm10.606 10.607h-7.07l2.828-2.829l-3.535-3.536l1.414-1.414l3.535 3.536l2.828-2.829v7.072z" fill="currentColor"/>`),
		g.Group(children),
	)
}