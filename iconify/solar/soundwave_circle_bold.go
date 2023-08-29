package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundwaveCircleBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12 2C6.477 2 2 6.477 2 12s4.477 10 10 10s10-4.477 10-10S17.523 2 12 2Zm.75 5a.75.75 0 0 0-1.5 0v10a.75.75 0 0 0 1.5 0V7Zm-5 2a.75.75 0 0 0-1.5 0v6a.75.75 0 0 0 1.5 0V9Zm10 1a.75.75 0 0 0-1.5 0v4a.75.75 0 0 0 1.5 0v-4Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}