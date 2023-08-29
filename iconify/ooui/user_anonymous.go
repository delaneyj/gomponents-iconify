package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UserAnonymous(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M15 2H5L4 8h12zM0 10s2 1 10 1s10-1 10-1l-4-2H4zm8 4h4v1H8z"/><circle cx="6" cy="15" r="3" fill="currentColor"/><circle cx="14" cy="15" r="3" fill="currentColor"/>`),
		g.Group(children),
	)
}