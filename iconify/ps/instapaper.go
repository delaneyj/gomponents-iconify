package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Instapaper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M209 449q-42-3-54-12q-13-10-13-50V77q0-39 13-49q11-11 54-13V2H4v13q43 1 55 13q13 10 13 49v310q0 40-13 50q-12 9-55 12v13h205v-13z"/>`),
		g.Group(children),
	)
}