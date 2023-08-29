package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SelectOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M12 6H8V2H6v4H2v2h4v4h2V8h4V6zm4 0h4v2h-4zm8 0v2h4v4h2V8a2 2 0 0 0-2-2zM6 16h2v4H6zm2 12v-4H6v4a2 2 0 0 0 2 2h4v-2zm20-12h2v4h-2zM16 28h4v2h-4zm12-4v4h-4v2h4a2 2 0 0 0 2-2v-4z"/>`),
		g.Group(children),
	)
}