package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BehanceLogoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M117.82 121.39A42 42 0 0 0 86 52H32a12 12 0 0 0-12 12v128a12 12 0 0 0 12 12h58a46 46 0 0 0 27.82-82.61ZM44 76h42a18 18 0 0 1 0 36H44Zm46 104H44v-44h46a22 22 0 0 1 0 44Zm66-104a12 12 0 0 1 12-12h64a12 12 0 0 1 0 24h-64a12 12 0 0 1-12-12Zm44 24a52 52 0 0 0 0 104a51.45 51.45 0 0 0 22.7-5.21a12 12 0 1 0-10.49-21.58A27.73 27.73 0 0 1 200 180a28.05 28.05 0 0 1-25.3-16H240a12 12 0 0 0 12-12a52.06 52.06 0 0 0-52-52Zm-25.3 40a28 28 0 0 1 50.6 0Z"/>`),
		g.Group(children),
	)
}