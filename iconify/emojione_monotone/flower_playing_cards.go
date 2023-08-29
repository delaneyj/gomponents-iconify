package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlowerPlayingCards(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M50.96 2H13.039C10.812 2 9 3.724 9 5.843v52.313C9 60.275 10.812 62 13.039 62H50.96c2.229 0 4.04-1.725 4.04-3.844V5.843C55 3.724 53.188 2 50.96 2zM10.84 7.435c0-1.988 1.296-3.56 3.384-3.56h35.553c2.088 0 3.384 1.762 3.384 3.75v45.846c-2.933-11.4-22.021-19.416-42.226-19.545l-.095-.051V7.435z"/><ellipse cx="25.519" cy="19.771" fill="currentColor" rx="12.002" ry="12.229"/>`),
		g.Group(children),
	)
}