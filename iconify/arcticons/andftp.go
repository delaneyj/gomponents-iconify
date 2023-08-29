package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Andftp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 38.02V22.18h15.84v15.84H4.5Zm23.16 0V22.18h15.83v15.84H27.66Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.93 34.74v-9.3h9.3v9.3h-9.3Zm-23.16 0v-9.3h9.3v9.3h-9.3Zm12.57-2.82h7.31m.01-2.44h-2.44V12.42h4.87V9.98H17.9v2.44h4.87v17.06h-2.44M6.94 11.2h10.97m12.18 0h10.97"/>`),
		g.Group(children),
	)
}