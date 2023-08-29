package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gym(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M9 6a3 3 0 0 0-3 3v30a3 3 0 0 0 3 3h30a3 3 0 0 0 3-3V9a3 3 0 0 0-3-3H9Zm23 10h-3v16h3V16Zm2 3h3v4h1v2h-1v4h-3V19ZM16 32h3V16h-3v16Zm-2-3h-3v-4h-1v-2h1v-4h3v10Zm7-4h6v-2h-6v2Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}