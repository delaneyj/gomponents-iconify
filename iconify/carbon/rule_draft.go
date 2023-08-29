package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RuleDraft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m29.707 19.293l-3-3a1 1 0 0 0-1.414 0L16 25.586V30h4.414l9.293-9.293a1 1 0 0 0 0-1.414zM19.586 28H18v-1.586l5-5L24.586 23zM26 21.586L24.414 20L26 18.414L27.586 20zM8 16h10v2H8zm0-6h12v2H8z"/><path fill="currentColor" d="M26 4a2.002 2.002 0 0 0-2-2H4a2.002 2.002 0 0 0-2 2v13a10.981 10.981 0 0 0 5.824 9.707L13 29.467V27.2l-4.234-2.258A8.986 8.986 0 0 1 4 17V4h20v9h2Z"/>`),
		g.Group(children),
	)
}