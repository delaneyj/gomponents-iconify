package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M30 19h-5v-5h-2v5h-5v2h5v5h2v-5h5v-2Z"/><path fill-rule="evenodd" d="M9 10a5 5 0 0 1 5-5h24a1 1 0 0 1 1 1v30a1 1 0 0 1-1 1H14a3 3 0 0 0-3 3v1h27v2H10a1 1 0 0 1-1-1V10Zm5 25a4.978 4.978 0 0 0-3 1V10a3 3 0 0 1 3-3h23v28H14Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}