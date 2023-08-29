package subway

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloth(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="m0 512l149.3-42.7V0L0 42.7V512zM362.7 42.7V512L512 469.3V0L362.7 42.7zm-192 426.6L341.3 512V42.7L170.7 0v469.3z"/>`),
		g.Group(children),
	)
}