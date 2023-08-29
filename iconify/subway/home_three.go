package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HomeThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M257 0L1 256l42.7 42.7L257 85.3l213.3 213.3L513 256L257 0zM107.7 298.7V512h298.7V298.7L257 149.3L107.7 298.7z"/>`),
		g.Group(children),
	)
}