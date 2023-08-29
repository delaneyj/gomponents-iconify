package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListBulletLtr(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M7 15h12v2H7zm0-6h12v2H7zm0-6h12v2H7z"/><circle cx="3" cy="4" r="2" fill="currentColor"/><circle cx="3" cy="10" r="2" fill="currentColor"/><circle cx="3" cy="16" r="2" fill="currentColor"/>`),
		g.Group(children),
	)
}