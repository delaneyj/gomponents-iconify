package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 19.937a1.243 1.243 0 0 1-.833-.319l-1.886-1.686a1.531 1.531 0 0 0-1.08-.458h-3.64a2.5 2.5 0 0 1-2.5-2.5l.006-8.41a2.5 2.5 0 0 1 2.5-2.5h14.872a2.5 2.5 0 0 1 2.5 2.5v8.411a2.5 2.5 0 0 1-2.5 2.5H15.79a1.483 1.483 0 0 0-1.062.441l-1.895 1.7a1.243 1.243 0 0 1-.833.321ZM4.567 5.063a1.5 1.5 0 0 0-1.5 1.5l-.006 8.411a1.5 1.5 0 0 0 1.5 1.5H8.2a2.483 2.483 0 0 1 1.767.732l1.864 1.667a.248.248 0 0 0 .333 0l1.874-1.682a2.5 2.5 0 0 1 1.751-.716h3.649a1.5 1.5 0 0 0 1.5-1.5V6.563a1.5 1.5 0 0 0-1.5-1.5Z"/>`),
		g.Group(children),
	)
}