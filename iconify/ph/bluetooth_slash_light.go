package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothSlashLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="m212.44 212l-160-176a6 6 0 0 0-8.88 8l72.32 79.55L52.4 171.2a6 6 0 0 0 7.2 9.6L114 140v84a6 6 0 0 0 9.6 4.8l52.28-39.21L203.56 220a6 6 0 0 0 8.88-8.08ZM126 212v-72l15.09 11.31l26.68 29.36ZM114 71.63V32a6 6 0 0 1 9.6-4.8l64 48a6 6 0 0 1 0 9.6L154.07 110a6 6 0 0 1-7.2-9.6L174 80l-48-36v27.63a6 6 0 0 1-12 0Z"/>`),
		g.Group(children),
	)
}