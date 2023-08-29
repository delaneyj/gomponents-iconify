package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Libremmail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M41.94 10.66H6.06a1.56 1.56 0 0 0-1.56 1.56v23.56a1.56 1.56 0 0 0 1.56 1.56h35.88a1.56 1.56 0 0 0 1.56-1.56V12.22a1.56 1.56 0 0 0-1.56-1.56Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.48 14.22L20.2 25.94a5.31 5.31 0 0 0 7.52 0l11.77-11.77m-.23 19.66l-10.63-8.71M8.76 33.83l10.63-8.71"/>`),
		g.Group(children),
	)
}