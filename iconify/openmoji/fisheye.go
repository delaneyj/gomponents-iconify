package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fisheye(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36" cy="36" r="28" fill="#61B2E4"/><circle r="18" fill="#92D3F5" transform="matrix(-1 0 0 1 36 36)"/><path fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="10" d="M54.125 36c0-2.118-.363-4.152-1.031-6.042c-2.489-7.04-9.202-12.083-17.094-12.083"/><circle cx="36" cy="36" r="7.25"/><circle cx="36" cy="36" r="28" fill="none" stroke="#000" stroke-width="2"/>`),
		g.Group(children),
	)
}