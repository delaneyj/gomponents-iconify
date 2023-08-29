package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MigrateFlasher(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M4.5 41.03L43.5 24L4.5 6.97m39.082 21.067l-9.143 4.318l6.729 5.047l-9.729 3.628"/>`),
		g.Group(children),
	)
}