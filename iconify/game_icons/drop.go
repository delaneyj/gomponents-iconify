package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M263.844 40.344C234.1 213.202 145.594 248.03 145.594 369.22c0 60.804 60.106 105.5 118.25 105.5c59.45 0 115.937-41.803 115.937-99.533c0-116.332-85.2-162.312-115.936-334.843zm-58.28 217.094c-27.963 75.53-5.105 154.567 54.25 179.375c15.185 6.348 31.724 7.714 47.905 6.28c-116.134 49.787-185.836-79.816-102.158-185.656z"/>`),
		g.Group(children),
	)
}