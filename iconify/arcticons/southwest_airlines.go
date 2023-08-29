package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SouthwestAirlines(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 41.277s12.33-8.16 16.16-16.426c.714-1.862 2.191-4.776 1.52-9.372c-.59-4.055-3.272-7.114-7.504-7.894C28.112 6.468 24 9.74 24 9.74s-4.112-3.272-10.175-2.155c-4.233.78-6.914 3.84-7.506 7.894c-.67 4.596.807 7.51 1.522 9.372C11.67 33.117 24 41.277 24 41.277h0Zm0-31.538l17.168 12.424M6.522 14.44l25.376 20.614"/>`),
		g.Group(children),
	)
}