package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Canonical(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M24 12c0 6.627-5.373 12-12 12c-6.628 0-12-5.373-12-12C0 5.372 5.372 0 12 0c6.627 0 12 5.372 12 12zM12 2.92A9.08 9.08 0 0 0 2.92 12A9.08 9.08 0 0 0 12 21.08A9.08 9.08 0 0 0 21.081 12A9.08 9.08 0 0 0 12 2.92zm0 16.722A7.64 7.64 0 0 1 4.36 12A7.64 7.64 0 0 1 12 4.36A7.64 7.64 0 0 1 19.641 12a7.64 7.64 0 0 1-7.64 7.641z"/>`),
		g.Group(children),
	)
}