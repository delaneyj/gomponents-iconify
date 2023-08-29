package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MobileAudio(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M28 12h2v3h-2zm-4-5h2v13h-2zm-4 4h2v5h-2zm-4-2h2v9h-2zm-4 3h2v3h-2z"/><circle cx="13.5" cy="24.5" r="1.5" fill="currentColor"/><path fill="currentColor" d="M20 30H7a2.002 2.002 0 0 1-2-2V4a2.002 2.002 0 0 1 2-2h13v2H7v24h13v-4h2v4a2.002 2.002 0 0 1-2 2Z"/>`),
		g.Group(children),
	)
}