package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceLaugh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 512a256 256 0 1 0 0-512a256 256 0 1 0 0 512zM96.8 314.1c-3.8-13.7 7.4-26.1 21.6-26.1h275.2c14.2 0 25.5 12.4 21.6 26.1C396.2 382 332.1 432 256 432S115.8 382 96.8 314.1zM144.4 192a32 32 0 1 1 64 0a32 32 0 1 1-64 0zm192-32a32 32 0 1 1 0 64a32 32 0 1 1 0-64z"/>`),
		g.Group(children),
	)
}