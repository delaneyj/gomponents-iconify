package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lockers(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M192 72a8 8 0 0 1-8 8h-24a8 8 0 0 1 0-16h24a8 8 0 0 1 8 8Zm-8 24h-24a8 8 0 0 0 0 16h24a8 8 0 0 0 0-16ZM72 80h24a8 8 0 0 0 0-16H72a8 8 0 0 0 0 16Zm24 16H72a8 8 0 0 0 0 16h24a8 8 0 0 0 0-16Zm128-48v176a8 8 0 0 1-16 0v-16h-72v16a8 8 0 0 1-16 0v-16H48v16a8 8 0 0 1-16 0V48a16 16 0 0 1 16-16h160a16 16 0 0 1 16 16ZM120 192V48H48v144Zm16 0h72V48h-72Z"/>`),
		g.Group(children),
	)
}