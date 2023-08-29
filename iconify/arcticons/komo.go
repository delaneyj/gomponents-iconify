package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Komo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.272 4.5v17.313L14.46 24l-2.188 2.187V43.5m23.456-39L16.647 24l19.081 19.5"/>`),
		g.Group(children),
	)
}