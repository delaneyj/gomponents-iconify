package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NumberCircleFiveTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M10.748 7a.75.75 0 0 0-.748.697l-.25 3.527a.75.75 0 0 0 .791.801h.015l.047-.003a26.088 26.088 0 0 1 .743-.028c.462-.011.936-.008 1.184.026a2 2 0 1 1-2.097 2.816a.75.75 0 0 0-1.362.628a3.5 3.5 0 1 0 3.667-4.93h-.003c-.385-.053-.978-.05-1.425-.04h-.005l.142-1.994h2.802a.75.75 0 0 0 0-1.5h-3.501ZM12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2ZM3.5 12a8.5 8.5 0 1 1 17 0a8.5 8.5 0 0 1-17 0Z"/>`),
		g.Group(children),
	)
}