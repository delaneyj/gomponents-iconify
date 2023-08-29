package octicon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DotTwentyFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M12 18a6 6 0 1 1 0-12a6 6 0 0 1 0 12Zm0-1.5a4.5 4.5 0 1 0 0-9a4.5 4.5 0 0 0 0 9Z"/>`),
		g.Group(children),
	)
}