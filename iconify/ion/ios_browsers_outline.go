package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IosBrowsersOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path d="M64 144v304h303.9V144H64zm287.9 288H80V160h271.9v272z" fill="currentColor"/><path d="M448 64H144v64h16V80h272v272h-48v16h64z" fill="currentColor"/>`),
		g.Group(children),
	)
}