package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Certilia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.383 24c0-11.787 12.557-19.5 21.705-19.5c2.395 0 4.974.57 7.529 1.458l-5.376 9.006c-6.1-.18-19.552-1.48-23.857 9.036Zm0 0c0 11.787 12.557 19.5 21.705 19.5c2.395 0 4.974-.57 7.529-1.458c0 0-15.421-3.938-19.906-18.042H9.384Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 14.985L18.711 24H9.383"/>`),
		g.Group(children),
	)
}