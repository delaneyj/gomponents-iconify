package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Redmoi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.544 30.004V17.996h2.702A5.254 5.254 0 0 1 36.5 23.25v1.501a5.254 5.254 0 0 1-5.254 5.254Zm-17.044 0V17.996h3.931a4.033 4.033 0 0 1 0 8.066H11.5m3.931 0l3.931 3.939m9.127-3.139h-5.254m5.254-5.724h-5.254"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2Z"/>`),
		g.Group(children),
	)
}