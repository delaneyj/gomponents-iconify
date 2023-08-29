package noto_v_1

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StopButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 128 128"),
		g.Raw(`<path fill="none" d="M0 0h128v128H0V0z"/><path fill="#40c0e7" d="M32 32h64v64H32V32z"/>`),
		g.Group(children),
	)
}