package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tecka(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.5 20.866h4.629m-2.27 6.988v-6.987m5.67 6.987h3.493m-3.493-6.988h3.493m-3.493 3.494h2.27m-2.27-3.494v6.987m11.901-2.442c0 1.31-1.048 2.358-2.358 2.358s-2.358-1.048-2.358-2.358v-2.358c0-1.31 1.048-2.358 2.358-2.358s2.27 1.048 2.27 2.358m3.72-2.327v6.987m.874-3.493h-.873M41.5 27.842l-2.27-6.987l-2.359 6.987m.787-2.359h3.056m-6.896 2.23l-2.882-3.493l2.882-3.493m-9.014-2.581l-.786.437l-.786-.437"/>`),
		g.Group(children),
	)
}