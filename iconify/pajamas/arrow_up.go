package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowUp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M10.72 5.841a.75.75 0 1 0 1.06-1.06L8.53 1.53L8 1l-.53.53l-3.25 3.25a.75.75 0 0 0 1.06 1.061l1.97-1.97V14.25a.75.75 0 0 0 1.5 0V3.871l1.97 1.97Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}