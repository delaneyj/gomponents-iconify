package material_symbols

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bedtime(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13.1 23q-2.1 0-3.937-.8t-3.2-2.163Q4.6 18.675 3.8 16.837T3 12.9q0-3.65 2.325-6.438T11.25 3q-.45 2.475.275 4.838t2.5 4.137q1.775 1.775 4.138 2.5T23 14.75q-.65 3.6-3.45 5.925T13.1 23Z"/>`),
		g.Group(children),
	)
}