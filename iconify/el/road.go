package el

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Road(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1200 1200"),
		g.Raw(`<path fill="currentColor" d="M1200 1200L957.743 0H658.691l9.164 276.675H532.144L541.308 0H242.256L0 1200h501.562l11.441-345.445h173.992L698.438 1200H1200M683.573 751.231H516.426l13.479-406.965h140.188l13.48 406.965"/>`),
		g.Group(children),
	)
}