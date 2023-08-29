package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RecordVinyl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0zm256-96a96 96 0 1 1 0 192a96 96 0 1 1 0-192zm0 224a128 128 0 1 0 0-256a128 128 0 1 0 0 256zm0-96a32 32 0 1 0 0-64a32 32 0 1 0 0 64z"/>`),
		g.Group(children),
	)
}