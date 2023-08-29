package ic

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SharpStyle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M3.87 20.21v-9.03l-3.19 7.7l3.19 1.33zm18.92-2.43L16.31 2.14L5.26 6.71l6.48 15.64l11.05-4.57zM7.88 8.75c-.55 0-1-.45-1-1s.45-1 1-1s1 .45 1 1s-.45 1-1 1zm-2 13h3.45l-3.45-8.34v8.34z"/>`),
		g.Group(children),
	)
}