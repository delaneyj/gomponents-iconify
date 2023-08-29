package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Horizontal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M330 129H43v179c-15 4-31 11-43 21V68C0 31 31 0 68 0h237c37 0 68 31 68 68v238h-43V129zm141 84h69c-3-28-15-54-35-74c-27-27-62-38-96-35V53c47-3 96 13 132 50c31 31 47 69 50 110h67l-92 100zM68 344h552c37 0 67 31 67 68v238c0 37-30 67-67 67H68c-37 0-68-30-68-67V412c0-37 31-68 68-68zm66 329h424V387H134v286zM78 568c20 0 38-17 38-37s-18-37-38-37s-37 17-37 37s17 37 37 37z"/>`),
		g.Group(children),
	)
}