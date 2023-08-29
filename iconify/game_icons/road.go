package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Road(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M168.4 29.55L24.38 477.6l15.24 4.8L183.6 34.45l-15.2-4.9zm175.2 0l-15.2 4.9l144 447.95l15.2-4.8l-144-448.05zM248 32l-.8 20h17.6l-.8-20h-16zm-1.4 36l-.8 20h20.4l-.8-20h-18.8zm-1.5 36l-1.3 32h24.4l-1.3-32h-21.8zm-1.9 48l-2 48h29.6l-2-48h-25.6zm-2.8 68l-2.4 60h36l-2.4-60h-31.2zm-3.3 84l-2.9 72h43.6l-2.9-72h-37.8zm-4 100l-3.1 76h52l-3.1-76h-45.8z"/>`),
		g.Group(children),
	)
}