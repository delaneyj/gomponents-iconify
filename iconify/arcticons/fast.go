package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fast(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.193 32.609H43.5a19.5 19.5 0 0 0-39 0h3.308"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.914 32.609A13.914 13.914 0 0 0 24 18.695m7.633 13.914A7.633 7.633 0 0 0 24 24.976"/><circle cx="24" cy="32.609" r="2.282" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 30.327V14.786"/>`),
		g.Group(children),
	)
}