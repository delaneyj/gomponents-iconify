package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RemoteControlTwoFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 2a1 1 0 0 1 1 1v18a1 1 0 0 1-1 1H6a1 1 0 0 1-1-1V3a1 1 0 0 1 1-1h12Zm-3 13h-2v2h2v-2Zm-4 0H9v2h2v-2Zm2-9h-2v2H9v2h1.999L11 12h2l-.001-2H15V8h-2V6Z"/>`),
		g.Group(children),
	)
}