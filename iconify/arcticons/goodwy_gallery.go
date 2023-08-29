package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GoodwyGallery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="24" cy="24" r="6.959"/><circle cx="24" cy="11.373" r="8.873"/><circle cx="32.929" cy="15.071" r="8.873"/><circle cx="36.627" cy="24" r="8.873"/><circle cx="32.929" cy="32.929" r="8.873"/><circle cx="24" cy="36.627" r="8.873"/><circle cx="15.071" cy="32.929" r="8.873"/><circle cx="11.373" cy="24" r="8.873"/><circle cx="15.071" cy="15.071" r="8.873"/></g>`),
		g.Group(children),
	)
}