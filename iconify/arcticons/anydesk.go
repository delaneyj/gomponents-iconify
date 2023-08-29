package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Anydesk(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.78 11.1L4.45 24l13.33 13.32L31.11 24Zm5.82 20.4l6.66 6.1L43.46 24l-13.2-12.9l-6.41 5.9"/>`),
		g.Group(children),
	)
}