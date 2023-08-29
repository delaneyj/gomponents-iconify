package uit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowCircleUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.354 7.147a.5.5 0 0 0-.344-.145L12 7l-.007.001a.5.5 0 0 0-.347.146l-4.5 4.5a.5.5 0 0 0 .708.707L11.5 8.707v7.794a.5.5 0 1 0 1 0V8.706l3.646 3.647a.498.498 0 0 0 .707 0a.5.5 0 0 0 0-.707l-4.5-4.5zM12 2C6.477 2 2 6.477 2 12s4.477 10 10 10c5.52-.006 9.994-4.48 10-10c0-5.523-4.477-10-10-10zm0 19a9 9 0 1 1 0-18a9.01 9.01 0 0 1 9 9a9 9 0 0 1-9 9z"/>`),
		g.Group(children),
	)
}