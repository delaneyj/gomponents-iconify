package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Airplay(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M22.996 30H9.004a1.002 1.002 0 0 1-.821-1.577l6.998-9.996a1 1 0 0 1 1.638 0l6.998 9.996a1.002 1.002 0 0 1-.82 1.577ZM10.92 28h10.16L16 20.744Z"/><path fill="currentColor" d="M28 24h-4v-2h4V6H4v16h4v2H4a2.002 2.002 0 0 1-2-2V6a2.002 2.002 0 0 1 2-2h24a2.002 2.002 0 0 1 2 2v16a2.002 2.002 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}