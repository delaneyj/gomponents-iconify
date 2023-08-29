package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BearSmileLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 17a4 4 0 0 0 4-4h-2a2 2 0 1 1-4 0H8a4 4 0 0 0 4 4ZM6.5 2a4.5 4.5 0 0 0-2.95 7.898a9 9 0 1 0 16.901 0a4.5 4.5 0 1 0-6.79-5.745a9.044 9.044 0 0 0-3.321 0A4.496 4.496 0 0 0 6.5 2ZM4 6.5a2.5 2.5 0 0 1 4.852-.851l.318.878l.898-.257A7.004 7.004 0 0 1 12 6c.672 0 1.32.094 1.932.27l.898.257l.318-.878a2.501 2.501 0 1 1 3.58 3.03l-.814.46l.404.842a7 7 0 1 1-12.635 0l.403-.843l-.814-.46A2.499 2.499 0 0 1 4 6.5Z"/>`),
		g.Group(children),
	)
}