package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Motonotifications(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.2 12.59H15.88L4.5 8.1v28.44a3.32 3.32 0 0 0 3.3 3.32h32.4a3.32 3.32 0 0 0 3.3-3.32V15.89a3.3 3.3 0 0 0-3.3-3.3Zm-24.89 18a3.12 3.12 0 1 1-3.11-3.13a3.13 3.13 0 0 1 3.11 3.13Zm0-9a3.12 3.12 0 1 1-3.11-3.13a3.13 3.13 0 0 1 3.11 3.13Zm2.61 8.93h20.82m-20.82-8.96h20.82"/>`),
		g.Group(children),
	)
}