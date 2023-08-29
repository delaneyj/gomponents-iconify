package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ParachuteFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M232 120A104.12 104.12 0 0 0 128 16A104.12 104.12 0 0 0 24 120a8 8 0 0 0 3.21 6.39L120 196v20h-8a8 8 0 0 0 0 16h32a8 8 0 0 0 0-16h-8v-20l92.8-69.6a8 8 0 0 0 3.2-6.4Zm-16.36-8h-39.81c-1.54-37.95-13.91-62.43-25.11-77a88.2 88.2 0 0 1 64.92 77Zm-61.24 16L128 175.53L101.6 128Zm-71.11 0l19.5 35.09L56 128Zm89.42 0H200l-46.79 35.09Zm-67.43-93c-11.2 14.57-23.57 39.05-25.11 77H40.36a88.2 88.2 0 0 1 64.92-77Z"/>`),
		g.Group(children),
	)
}