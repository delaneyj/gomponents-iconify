package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gitter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width=".87" d="M12.65 4.5v22.34m7.57-14.75V43.5m7.56 0V12.09m7.57 0v14.75"/>`),
		g.Group(children),
	)
}