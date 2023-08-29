package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ProhibitedTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12 2c5.523 0 10 4.477 10 10s-4.477 10-10 10S2 17.523 2 12S6.477 2 12 2Zm6.517 4.543L6.543 18.517A8.5 8.5 0 0 0 18.517 6.543ZM12 3.5a8.5 8.5 0 0 0-6.517 13.957L17.457 5.483A8.466 8.466 0 0 0 12 3.5Z"/>`),
		g.Group(children),
	)
}