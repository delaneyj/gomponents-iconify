package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dictionaryformids(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m22.96 28l-2.6-8l-2.7 8m.9-2.7h3.5m7.31-5.3h5.3l-5.3 8h5.3m-10.505-4h4M8.4 6.5v35a2 2 0 0 0 2 2h2.33v-39H10.4a2 2 0 0 0-2 2Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.73 4.5v39H37.6a2 2 0 0 0 2-2v-35a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}