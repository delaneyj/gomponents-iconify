package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Nut(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M448 1024L0 768V256L448 0l448 256v512zm.5-704q-79.5 0-136 56T256 512t56 136t136 56t136-56t56-136t-56-136t-135.5-56z"/>`),
		g.Group(children),
	)
}