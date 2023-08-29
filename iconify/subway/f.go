package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func F(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M69.8 0C31.3 0 0 31.2 0 69.8c0 38.6 31.3 69.8 69.8 69.8c38.6 0 69.8-31.3 69.8-69.8C139.6 31.2 108.4 0 69.8 0zm0 104.7c-19.3 0-34.9-15.6-34.9-34.9s15.6-34.9 34.9-34.9c19.3 0 34.9 15.6 34.9 34.9s-15.6 34.9-34.9 34.9zM512 93.1V0H209.5v512h116.4V302.5h162.9v-93.1h-163V93.1H512z"/>`),
		g.Group(children),
	)
}