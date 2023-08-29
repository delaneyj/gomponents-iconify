package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BrandAdobe(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m12.893 4.514l7.977 14a.993.993 0 0 1-.394 1.365a1.04 1.04 0 0 1-.5.127H16.5l-4.5-8l-2.5 4H11l2 4H4.023c-.565 0-1.023-.45-1.023-1c0-.171.045-.34.13-.49l7.977-13.993a1.034 1.034 0 0 1 1.786 0z"/>`),
		g.Group(children),
	)
}