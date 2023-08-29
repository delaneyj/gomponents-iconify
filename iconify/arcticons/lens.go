package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lens(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15 5.5A9.49 9.49 0 0 0 5.5 15v4.15a4.21 4.21 0 0 1 .5 0h4v-4.1A5.08 5.08 0 0 1 15.05 10h4.08V6a4.21 4.21 0 0 1 0-.5Zm13.81 0a4.21 4.21 0 0 1 0 .5v4H33a5.07 5.07 0 0 1 5 5.05v4.08h4a4.21 4.21 0 0 1 .5 0V15A9.49 9.49 0 0 0 33 5.5ZM24 16.3a7.68 7.68 0 1 0 5.45 2.25A7.7 7.7 0 0 0 24 16.3ZM5.5 28.83V33a9.49 9.49 0 0 0 9.5 9.5h4.15a4.21 4.21 0 0 1 0-.5v-4h-4.1A5.07 5.07 0 0 1 10 33v-4.13H6a4.21 4.21 0 0 1-.5 0Zm30.94 3.5a4.11 4.11 0 1 0 4.11 4.11a4.11 4.11 0 0 0-4.11-4.11Zm0 0"/>`),
		g.Group(children),
	)
}