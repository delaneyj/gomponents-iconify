package vscode_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileTypeSketch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="#fdb300" d="M8.109 4.26L16 3.433l7.891.828L30 12.4L16 28.567L2 12.4Z"/><path fill="#ea6c00" d="M7.671 12.395L16 28.567L2 12.395h5.671zm16.658 0L16 28.567l14-16.172h-5.671z"/><path fill="#fdad00" d="M7.671 12.395h16.658L16 28.567L7.671 12.395z"/><path fill="#fdd231" d="m16 3.433l-7.891.827l-.438 8.135L16 3.433zm0 0l7.891.827l.438 8.135L16 3.433z"/><path fill="#fdad00" d="M30 12.395L23.891 4.26l.438 8.135H30zm-28 0L8.109 4.26l-.438 8.135H2z"/><path fill="#feeeb7" d="m16 3.433l-8.329 8.962h16.658L16 3.433z"/>`),
		g.Group(children),
	)
}