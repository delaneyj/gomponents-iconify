package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ColorPicker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M16.308 4.434c.912-.911.81-2.495-.229-3.534c-1.041-1.039-2.624-1.142-3.534-.23l-1.181 1.183l3.764 3.764l1.18-1.183zm-9.84 9.842l6.147-6.148l.772.772l1.584-1.586l-5.309-5.309l-1.585 1.584l.773.773l-6.148 6.146l-1.682 4.693l.754.754l4.694-1.679zm3.448-8.848l1.635 1.636l-6.262 6.26l-2.594.96l.959-2.596l6.262-6.26z"/>`),
		g.Group(children),
	)
}