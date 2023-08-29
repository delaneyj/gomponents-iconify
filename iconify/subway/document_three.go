package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M372.4 512L512 372.4H372.4V512zM93.1 93.1V512h256V349.1H512v-256H93.1zM418.9 0H0v418.9h46.5V46.5h372.4V0z"/>`),
		g.Group(children),
	)
}