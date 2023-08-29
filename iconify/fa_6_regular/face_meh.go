package fa_6_regular

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceMeh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M464 256a208 208 0 1 1-416 0a208 208 0 1 1 416 0zM256 0a256 256 0 1 0 0 512a256 256 0 1 0 0-512zm-79.6 240a32 32 0 1 0 0-64a32 32 0 1 0 0 64zm192-32a32 32 0 1 0-64 0a32 32 0 1 0 64 0zM184 328c-13.3 0-24 10.7-24 24s10.7 24 24 24h144c13.3 0 24-10.7 24-24s-10.7-24-24-24H184z"/>`),
		g.Group(children),
	)
}