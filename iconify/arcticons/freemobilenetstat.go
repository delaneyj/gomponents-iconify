package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Freemobilenetstat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.42 35.89a16.81 16.81 0 0 1 23.77-23.78L21.31 24l11.88 11.89a16.82 16.82 0 0 1-23.77 0Zm29.15 0L26.69 24l11.88-11.89a16.8 16.8 0 0 1 0 23.78Zm0 0"/>`),
		g.Group(children),
	)
}