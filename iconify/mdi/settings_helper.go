package mdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SettingsHelper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M7 22h2v2H7v-2m4 0h2v2h-2v-2m4 0h2v2h-2v-2Z"/>`),
		g.Group(children),
	)
}