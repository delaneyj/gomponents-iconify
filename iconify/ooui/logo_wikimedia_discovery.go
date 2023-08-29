package ooui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LogoWikimediaDiscovery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 17c0 1.1-2 2-2 2s-2-.9-2-2m2-10a1.54 1.54 0 0 1-1.5-1.5a1.5 1.5 0 0 1 3 0A1.54 1.54 0 0 1 10 7zm3.3 4.7C14.1 7.9 12.7 1 10 1S5.8 7.7 6.6 11.5L5 15h2.7l.3 1h4c.2-.3.1-.5.3-1H15z"/>`),
		g.Group(children),
	)
}