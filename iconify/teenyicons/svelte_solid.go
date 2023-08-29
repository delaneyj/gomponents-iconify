package teenyicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SvelteSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M13.283 1.944a4.001 4.001 0 0 1-.27 4.622a4 4 0 0 1-1.503 5.09l-4.24 2.65A4 4 0 0 1 2.028 8.41a4 4 0 0 1 1.503-5.09L7.77.671a4 4 0 0 1 5.512 1.273ZM8.3 1.52a3 3 0 0 1 4.13 4.143a3.993 3.993 0 0 0-2.386-1.345l.724-.453l-.53-.848l-5.088 3.18l.53.847L7.8 5.72a3 3 0 1 1 3.18 5.088l-4.24 2.65a3 3 0 0 1-4.13-4.143a3.993 3.993 0 0 0 2.386 1.345l-.725.452l.53.848L9.89 8.78l-.53-.847l-2.12 1.325a3 3 0 0 1-3.18-5.089L8.3 1.52Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}