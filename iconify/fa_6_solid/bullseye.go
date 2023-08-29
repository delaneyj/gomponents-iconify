package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bullseye(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M448 256a192 192 0 1 0-384 0a192 192 0 1 0 384 0zM0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0zm256 80a80 80 0 1 0 0-160a80 80 0 1 0 0 160zm0-224a144 144 0 1 1 0 288a144 144 0 1 1 0-288zm-32 144a32 32 0 1 1 64 0a32 32 0 1 1-64 0z"/>`),
		g.Group(children),
	)
}