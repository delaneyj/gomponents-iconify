package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Adobe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20.136 2.667H32v26.667zm-8.272 0H0v26.667zM16 12.531l7.469 16.803h-5.068l-2.136-5.333h-5.463z"/>`),
		g.Group(children),
	)
}