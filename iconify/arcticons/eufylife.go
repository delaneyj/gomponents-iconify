package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eufylife(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M45.5 24A21.5 21.5 0 1 1 24 2.5V24"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.523 2.535a4.125 4.125 0 0 0-3.204 6.73h0l6.615 7.764l6.55-7.687l.033-.037l.032-.04h0a4.129 4.129 0 1 0-6.615-4.924a4.117 4.117 0 0 0-3.407-1.806Z"/>`),
		g.Group(children),
	)
}