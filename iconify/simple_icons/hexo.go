package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hexo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.02 0L1.596 6.02l-.02 12L11.978 24l10.426-6.02l.02-12zm4.828 17.14l-.96.558l-.969-.574V12.99H9.081v4.15l-.96.558l-.969-.574V6.854l.964-.552l.965.563v4.145h5.838V6.86l.965-.552l.964.563z"/>`),
		g.Group(children),
	)
}