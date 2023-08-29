package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SpeakerDeck(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M8 8c-2.757 0-5 2.243-5 5s2.243 5 5 5h6a1.001 1.001 0 0 1 0 2H5a2 2 0 0 0 0 4h9c2.757 0 5-2.243 5-5s-2.243-5-5-5H8a1.001 1.001 0 0 1 0-2h7a2 2 0 0 0 0-4H8zm10.445 0c.344.59.555 1.268.555 2c0 .734-.217 1.409-.56 2H24a1 1 0 0 1 1 1v6a1 1 0 0 1-1 1h-3.08a6.978 6.978 0 0 1-2.031 4H25a4 4 0 0 0 4-4v-8a4 4 0 0 0-4-4h-6.555z"/>`),
		g.Group(children),
	)
}