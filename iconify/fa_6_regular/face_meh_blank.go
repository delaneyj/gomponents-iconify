package fa_6_regular

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceMehBlank(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 48a208 208 0 1 0 0 416a208 208 0 1 0 0-416zm256 208a256 256 0 1 1-512 0a256 256 0 1 1 512 0zm-367.6-48a32 32 0 1 1 64 0a32 32 0 1 1-64 0zm192-32a32 32 0 1 1 0 64a32 32 0 1 1 0-64z"/>`),
		g.Group(children),
	)
}