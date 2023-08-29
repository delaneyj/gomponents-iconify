package ls

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Surprise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 717 717"),
		g.Raw(`<path fill="currentColor" d="M358 0c198 0 359 155 359 349S556 698 358 698S0 543 0 349S160 0 358 0zM119 236c0 58 47 106 105 106s106-48 106-106s-48-105-106-105s-105 47-105 105zm374 106c58 0 105-48 105-106s-47-105-105-105s-106 47-106 105s48 106 106 106zM271 236c0 26-21 47-47 47s-47-21-47-47s21-47 47-47s47 21 47 47zm222-47c26 0 47 21 47 47s-21 47-47 47s-48-21-48-47s22-47 48-47zM316 579h84c51 0 94-42 94-94s-43-94-94-94h-84c-51 0-93 42-93 94s42 94 93 94z"/>`),
		g.Group(children),
	)
}