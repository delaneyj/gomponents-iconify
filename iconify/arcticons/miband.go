package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Miband(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-miterlimit="10" d="M24 8.5a5.453 5.453 0 0 1 5.45 5.456v20.088a5.45 5.45 0 1 1-10.9 0V13.956A5.453 5.453 0 0 1 24 8.5Z"/><path fill="none" stroke="currentColor" d="M31.83 33.519c0 3.203.063 6.041-1.315 8.384S25.553 44.5 24.722 44.5h-1.444c-.83 0-4.415-.254-5.793-2.597s-1.314-5.181-1.314-8.384V14.48c0-3.203-.064-6.041 1.314-8.384S22.448 3.5 23.278 3.5h1.444c.83 0 4.415.254 5.793 2.597s1.314 5.181 1.314 8.384V33.52Z"/>`),
		g.Group(children),
	)
}