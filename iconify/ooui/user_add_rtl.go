package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserAddRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<circle cx="11.5" cy="10.5" r="3.5" fill="currentColor"/><path fill="currentColor" d="M6 0v4h4v2H6v4H4V6H0V4h4V0zm6 15c4.6 0 7 2.69 7 4.23V20H5v-.77C5 17.69 7.4 15 12 15z"/>`),
		g.Group(children),
	)
}