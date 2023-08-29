package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ntfy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.719 7.784a4.483 4.483 0 0 0-4.49 4.496v20.185a4.528 4.528 0 0 0 .166 1.202L4.5 40.216l12.294-3.255H39.01a4.483 4.483 0 0 0 4.49-4.496V12.28a4.483 4.483 0 0 0-4.49-4.496Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.227 27.878l10.282-5.945l-10.282-5.945m13.139 13.094h10.96"/>`),
		g.Group(children),
	)
}