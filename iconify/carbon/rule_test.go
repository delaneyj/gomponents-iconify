package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RuleTest(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M23 30a7 7 0 1 1 7-7a7.008 7.008 0 0 1-7 7Zm0-12a5 5 0 1 0 5 5a5.005 5.005 0 0 0-5-5Z"/><path fill="currentColor" d="m26 24.586l-2-2V20h-2v3.414L24.586 26L26 24.586zM8 16h6v2H8zm0-6h12v2H8z"/><path fill="currentColor" d="M26 4a2.002 2.002 0 0 0-2-2H4a2.002 2.002 0 0 0-2 2v13a10.981 10.981 0 0 0 5.824 9.707L13 29.467V27.2l-4.234-2.258A8.986 8.986 0 0 1 4 17V4h20v9h2Z"/>`),
		g.Group(children),
	)
}