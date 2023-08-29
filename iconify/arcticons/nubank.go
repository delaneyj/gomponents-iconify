package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nubank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.55 33.49H4.5V17.11a2.6 2.6 0 0 1 2.6-2.6h10.36v19h4.85V15.25a5.2 5.2 0 0 0-5.19-5.19h-4a2.59 2.59 0 0 0-2.6 2.6Zm26.9-18.98h6v16.38a2.6 2.6 0 0 1-2.6 2.6H30.54v-19h-4.85v18.26a5.2 5.2 0 0 0 5.19 5.19h4a2.59 2.59 0 0 0 2.6-2.6Z"/>`),
		g.Group(children),
	)
}