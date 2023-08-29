package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Comdirect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.53 3.552a20.323 20.323 0 0 0-12.955 5.952a20.585 20.585 0 0 0 0 28.991a20.291 20.291 0 0 0 28.819 0l-6.551-6.59a11.068 11.068 0 0 1-15.717 0a11.228 11.228 0 0 1 0-15.81a11.068 11.068 0 0 1 15.717 0l6.55-6.59A20.308 20.308 0 0 0 25.53 3.552Z"/>`),
		g.Group(children),
	)
}