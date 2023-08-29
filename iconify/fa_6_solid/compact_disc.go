package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CompactDisc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0zm256 32a32 32 0 1 1 0-64a32 32 0 1 1 0 64zm-96-32a96 96 0 1 0 192 0a96 96 0 1 0-192 0zm-64-16c0-35 17.5-71.1 45.2-98.8S205 96 240 96c8.8 0 16-7.2 16-16s-7.2-16-16-16c-45.4 0-89.2 22.3-121.5 54.5S64 194.6 64 240c0 8.8 7.2 16 16 16s16-7.2 16-16z"/>`),
		g.Group(children),
	)
}