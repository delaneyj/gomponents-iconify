package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChartPieSliceBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M93.82 116.64a12 12 0 0 0 6.18-10.49V40.74a12 12 0 0 0-17-10.91A108.26 108.26 0 0 0 20 128c0 3.37.16 6.76.47 10.1a12 12 0 0 0 17.76 9.38ZM76 62.06v37l-31.19 17.3A84.39 84.39 0 0 1 76 62.06Zm145.88 12.6c-.11-.24-.24-.47-.37-.7s-.26-.43-.4-.63A108.05 108.05 0 0 0 128 20a12 12 0 0 0-12 12v89.53l-76.82 44.74a12 12 0 0 0-4.3 16.46A108 108 0 0 0 236 128a107.22 107.22 0 0 0-14.12-53.34ZM140 44.85a84.08 84.08 0 0 1 54 31.24l-54 31.46ZM128 212a84.47 84.47 0 0 1-65.57-31.5l71.42-41.6l.36-.2L206 96.87A84 84 0 0 1 128 212Z"/>`),
		g.Group(children),
	)
}