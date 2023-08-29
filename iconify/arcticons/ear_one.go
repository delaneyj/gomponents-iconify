package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EarOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><circle cx="24" cy="24" r="10" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.451 42.5a3.555 3.555 0 0 1 3.555-3.519H24a14.982 14.982 0 0 0 9.09-26.89L23.656 5.5M5.5 28.172l4.68 2.843a4.295 4.295 0 0 1-2.23 7.966H5.5"/>`),
		g.Group(children),
	)
}