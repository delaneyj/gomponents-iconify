package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PixelFilter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.5 5.5h-33a2 2 0 0 0-2 2v33a2 2 0 0 0 2 2h33a2 2 0 0 0 2-2v-33a2 2 0 0 0-2-2ZM5.5 24h37M24 5.5v37m0-27.75h18.5M33.25 5.5v37m4.625-37v37M28.625 24v18.5m-9.25-9.25v9.25m-9.25-9.25v9.25m4.625 0V24M5.5 33.25h37m-9.25 2.313h9.25m-9.25 4.625h9.25M24 28.625h18.5m-9.25-9.25h9.25m-9.25-9.25h9.25m-37 27.75h37m-2.312-4.625v9.25m-4.625-9.25v9.25"/>`),
		g.Group(children),
	)
}