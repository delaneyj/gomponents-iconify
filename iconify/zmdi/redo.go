package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Redo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="m360 162l76-77v192H244l78-77q-48-40-110-40q-56 0-100.5 33T50 277L0 261q22-68 80.5-111T212 107q84 0 148 55z"/>`),
		g.Group(children),
	)
}