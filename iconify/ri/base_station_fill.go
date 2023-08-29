package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BaseStationFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m12 13l6 9H6l6-9Zm-1.06-2.44a1.5 1.5 0 1 1 2.121-2.12a1.5 1.5 0 0 1-2.122 2.12ZM5.281 2.782l1.415 1.414a7.5 7.5 0 0 0 0 10.607l-1.415 1.414a9.5 9.5 0 0 1 0-13.435Zm13.436 0a9.5 9.5 0 0 1 0 13.435l-1.415-1.414a7.5 7.5 0 0 0 0-10.607l1.415-1.414ZM8.11 5.611l1.414 1.414a3.5 3.5 0 0 0 0 4.95L8.11 13.389a5.5 5.5 0 0 1 0-7.778Zm7.778 0a5.5 5.5 0 0 1 0 7.778l-1.414-1.414a3.5 3.5 0 0 0 0-4.95l1.414-1.414Z"/>`),
		g.Group(children),
	)
}