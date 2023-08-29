package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartCustom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m29.707 19.293l-3-3a1 1 0 0 0-1.414 0L16 25.586V30h4.414l9.293-9.293a1 1 0 0 0 0-1.414ZM19.586 28H18v-1.586l5-5L24.586 23ZM26 21.586L24.414 20L26 18.414L27.586 20ZM18 18v-4h2v4zm-4 0v-6h2v6zm-4 0V8h2v10z"/><path fill="currentColor" d="M12 30H6a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h16a2.002 2.002 0 0 1 2 2v10h-2V4H6v24h6Z"/>`),
		g.Group(children),
	)
}