package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Callrecorder(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.41 4.5a7.51 7.51 0 0 1 7.52 7.5v10a7.53 7.53 0 1 1-15 0V12a7.52 7.52 0 0 1 7.48-7.5ZM4 22.08h7.86a13.56 13.56 0 0 0 13.55 13.56h0v7.86A21.42 21.42 0 0 1 4 22.08Z"/>`),
		g.Group(children),
	)
}