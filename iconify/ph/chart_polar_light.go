package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPolarLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M128 26a102 102 0 1 0 102 102A102.12 102.12 0 0 0 128 26Zm89.8 96h-28.1A62.07 62.07 0 0 0 134 66.3V38.2a90.15 90.15 0 0 1 83.8 83.8Zm-95.8 0H78.37A50.09 50.09 0 0 1 122 78.37Zm0 12v43.63A50.09 50.09 0 0 1 78.37 134Zm12 0h43.63A50.09 50.09 0 0 1 134 177.63Zm0-12V78.37A50.09 50.09 0 0 1 177.63 122Zm-12-83.8v28.1A62.07 62.07 0 0 0 66.3 122H38.2A90.15 90.15 0 0 1 122 38.2ZM38.2 134h28.1a62.07 62.07 0 0 0 55.7 55.7v28.1A90.15 90.15 0 0 1 38.2 134Zm95.8 83.8v-28.1a62.07 62.07 0 0 0 55.7-55.7h28.1a90.15 90.15 0 0 1-83.8 83.8Z"/>`),
		g.Group(children),
	)
}