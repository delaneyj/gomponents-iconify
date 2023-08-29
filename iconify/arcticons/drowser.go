package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drowser(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="19.198" height="44.407" x="15.644" y="1.797" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="9.599" transform="rotate(45 25.243 24)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.03 30.787L18.456 17.213"/>`),
		g.Group(children),
	)
}