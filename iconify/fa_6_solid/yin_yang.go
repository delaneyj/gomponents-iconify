package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YinYang(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 64c53 0 96 43 96 96s-43 96-96 96s-96 43-96 96s43 96 96 96c-106 0-192-86-192-192S150 64 256 64zm0 448a256 256 0 1 0 0-512a256 256 0 1 0 0 512zm32-352a32 32 0 1 0-64 0a32 32 0 1 0 64 0zm-64 192a32 32 0 1 1 64 0a32 32 0 1 1-64 0z"/>`),
		g.Group(children),
	)
}