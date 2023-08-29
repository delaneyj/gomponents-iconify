package gridicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloud(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 9c-.009 0-.017.002-.025.003A6.496 6.496 0 0 0 5 9.5a6.5 6.5 0 0 0 .186 1.519C5.123 11.016 5.064 11 5 11a4 4 0 0 0-4 4c0 1.202.541 2.267 1.38 3h18.593C22.196 17.089 23 15.643 23 14a5 5 0 0 0-5-5z"/>`),
		g.Group(children),
	)
}