package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GaugeHigh(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 256a256 256 0 1 1 512 0a256 256 0 1 1-512 0zM288 96a32 32 0 1 0-64 0a32 32 0 1 0 64 0zm-32 320c35.3 0 64-28.7 64-64c0-17.4-6.9-33.1-18.1-44.6L366 161.7c5.3-12.1-.2-26.3-12.3-31.6s-26.3.2-31.6 12.3L257.9 288H256c-35.3 0-64 28.7-64 64s28.7 64 64 64zm-80-272a32 32 0 1 0-64 0a32 32 0 1 0 64 0zM96 288a32 32 0 1 0 0-64a32 32 0 1 0 0 64zm352-32a32 32 0 1 0-64 0a32 32 0 1 0 64 0z"/>`),
		g.Group(children),
	)
}