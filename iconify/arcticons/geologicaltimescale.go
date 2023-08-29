package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Geologicaltimescale(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/><path fill="none" stroke="currentColor" d="M5.5 15.809h18.547V5.5M5.5 38.85h18.547V15.809H42.5m-18.453 6.633H42.5M24.047 8.89H42.5m-37 16.842h37M24.047 38.85H42.5m-18.453-3.452H42.5"/>`),
		g.Group(children),
	)
}