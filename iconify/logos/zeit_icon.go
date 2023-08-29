package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ZeitIcon(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<defs><linearGradient x1="100.93%" y1="181.283%" x2="41.769%" y2="100%" id="ssvg-id-zeit-icona"><stop stop-color="#FFF" offset="0%"/><stop offset="100%"/></linearGradient></defs><path fill="url(#ssvg-id-zeit-icona)" d="M128 0l128 227.093H0z"/>`),
		g.Group(children),
	)
}