package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GaiaGps(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.01 41.29l19.49-9.757l-9.51-14.226l-2.148 1.21L24.01 6.71l-5.11 7.706l-2.124-1.161L4.5 31.558l19.51 9.731Zm0-34.58v34.58"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.01 16.664l5.387 2.767l-5.387 2.84l-5.308-2.816l5.308-2.79Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24 24.716l4.928-2.568l3.58 1.85L24 28.224L15.59 24l3.408-1.877L24 24.716Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m24.05 30.817l7.421-3.804l4.397 2.075l-11.819 6.15l-11.868-6.1l4.249-2.15l7.62 3.829"/>`),
		g.Group(children),
	)
}