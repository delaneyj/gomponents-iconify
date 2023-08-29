package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Kite(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m43.5 37.145l-26.732.072C10.639 32.721 4.5 28.214 4.5 23.851c.01-4.373 6.15-8.829 12.299-13.068H43.5c-6.324 4.373-12.648 8.746-12.648 13.14S37.176 32.73 43.5 37.145Zm-26.62.072l14.239-14.885"/>`),
		g.Group(children),
	)
}