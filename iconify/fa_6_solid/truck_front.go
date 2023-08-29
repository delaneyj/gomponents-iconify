package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TruckFront(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M0 80C0 35.8 35.8 0 80 0h352c44.2 0 80 35.8 80 80v288c0 26.2-12.6 49.4-32 64v48c0 17.7-14.3 32-32 32h-32c-17.7 0-32-14.3-32-32v-32H128v32c0 17.7-14.3 32-32 32H64c-17.7 0-32-14.3-32-32v-48c-19.4-14.6-32-37.8-32-64V80zm129.9 72.2L112 224h288l-17.9-71.8C378.5 138 365.7 128 351 128H161c-14.7 0-27.5 10-31 24.2zM128 320a32 32 0 1 0-64 0a32 32 0 1 0 64 0zm288 32a32 32 0 1 0 0-64a32 32 0 1 0 0 64z"/>`),
		g.Group(children),
	)
}