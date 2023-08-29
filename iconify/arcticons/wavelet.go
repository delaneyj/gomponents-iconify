package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wavelet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.5 23.268a20.575 20.575 0 0 1-4.592-4.725s-.865 21.363-5.857 21.363c-5.99 0-2.862-31.812-9.051-31.812s-3.061 31.812-9.051 31.812c-4.992 0-5.857-21.363-5.857-21.363A20.575 20.575 0 0 1 4.5 23.268"/>`),
		g.Group(children),
	)
}