package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ListBulletRtl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M1 15h12v2H1zm0-6h12v2H1zm0-6h12v2H1z"/><circle cx="17" cy="4" r="2" fill="currentColor"/><circle cx="17" cy="10" r="2" fill="currentColor"/><circle cx="17" cy="16" r="2" fill="currentColor"/>`),
		g.Group(children),
	)
}