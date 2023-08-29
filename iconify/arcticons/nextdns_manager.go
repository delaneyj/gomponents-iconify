package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NextdnsManager(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5c12.48-7.639 16.676-17.483 16.676-31.577C34.974 10.686 31.262 9.772 24 4.5c-7.262 5.272-10.974 6.186-16.676 7.423C7.324 26.017 11.52 35.861 24 43.5Zm11.877-32.746l-22.72 22.997"/>`),
		g.Group(children),
	)
}