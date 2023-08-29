package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceSurprise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 512a256 256 0 1 0 0-512a256 256 0 1 0 0 512zm-79.6-336a32 32 0 1 1 0 64a32 32 0 1 1 0-64zm128 32a32 32 0 1 1 64 0a32 32 0 1 1-64 0zM256 288a64 64 0 1 1 0 128a64 64 0 1 1 0-128z"/>`),
		g.Group(children),
	)
}