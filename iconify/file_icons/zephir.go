package file_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Zephir(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M235.652 21.305L0 312.405l189.096 178.29L512 314.857L235.652 21.305zm228.839 285.96L207.647 447.13l41.288-368.84L464.49 307.265zM218.084 89.714l-39.278 350.889L39.987 309.717L218.084 89.714z"/>`),
		g.Group(children),
	)
}