package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Trello(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<defs><linearGradient id="logosTrello0" x1="50%" x2="50%" y1="0%" y2="100%"><stop offset="0%" stop-color="#0091E6"/><stop offset="100%" stop-color="#0079BF"/></linearGradient></defs><rect width="256" height="256" fill="url(#logosTrello0)" rx="25"/><rect width="78.08" height="112" x="144.64" y="33.28" fill="#FFF" rx="12"/><rect width="78.08" height="176" x="33.28" y="33.28" fill="#FFF" rx="12"/>`),
		g.Group(children),
	)
}