package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Greenpass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="10.467" r="2.824" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m12.12 28.156l7.83 7.784l15.93-15.748"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.561 8.583h9.018a2.882 2.882 0 0 1 2.925 2.7V40.35a3.317 3.317 0 0 1-2.835 3.15H9.331a2.634 2.634 0 0 1-2.835-2.7V11.958c.066-1.67.842-3.29 2.475-3.42l9.441.045h.027A5.491 5.491 0 0 1 24 4.5a5.491 5.491 0 0 1 5.561 4.083Z"/>`),
		g.Group(children),
	)
}