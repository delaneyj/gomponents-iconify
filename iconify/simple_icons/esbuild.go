package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Esbuild(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 0A12 12 0 0 0 0 12a12 12 0 0 0 12 12a12 12 0 0 0 12-12A12 12 0 0 0 12 0zM6.718 5.282L13.436 12l-6.718 6.718l-2.036-2.036L9.364 12L4.682 7.318zm7.2 0L20.636 12l-6.718 6.718l-2.036-2.036L16.564 12l-4.682-4.682z"/>`),
		g.Group(children),
	)
}