package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Redo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M354.5 65.1H236.3l118.2 118.2H137.8C61.7 183.2 0 244.9 0 321.1c0 76.1 61.7 137.8 137.8 137.8v-78.8c-32.6 0-59.1-26.4-59.1-59.1c0-32.6 26.4-59.1 59.1-59.1h216.6L236.3 380.2h118.2L512 222.6L354.5 65.1z"/>`),
		g.Group(children),
	)
}