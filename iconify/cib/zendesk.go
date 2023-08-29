package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zendesk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M14.781 28.125H0l14.781-17.844zm17.219 0H17.219a7.392 7.392 0 0 1 14.782 0zm-14.781-6.406V3.875H32zM14.781 3.875a7.392 7.392 0 0 1-14.782 0z"/>`),
		g.Group(children),
	)
}