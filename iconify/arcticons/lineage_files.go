package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LineageFiles(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.878 13.042H27.846l-4.084-4.084H11.207l-4.085 4.084A2.622 2.622 0 0 0 4.5 15.664V36.42a2.622 2.622 0 0 0 2.622 2.622h33.756A2.622 2.622 0 0 0 43.5 36.42V15.664a2.622 2.622 0 0 0-2.622-2.622Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.24 26.04A19.265 19.265 0 0 0 8.665 37.732q0 .66.043 1.307"/>`),
		g.Group(children),
	)
}