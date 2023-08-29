package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Linktree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.126 20.084h31.748M12.775 8.859l22.45 22.45m-22.45 0l22.45-22.45M24 20.098V4.21m.002 25.942V43.5"/>`),
		g.Group(children),
	)
}