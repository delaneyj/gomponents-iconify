package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CellSignalSlashLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M86 152v48a6 6 0 0 1-12 0v-48a6 6 0 0 1 12 0Zm-46 34a6 6 0 0 0-6 6v8a6 6 0 0 0 12 0v-8a6 6 0 0 0-6-6Zm172.44 26l-160-176a6 6 0 0 0-8.88 8L114 121.52V200a6 6 0 0 0 12 0v-65.28l28 30.8V200a6 6 0 0 0 12 0v-21.28L203.56 220a6 6 0 0 0 8.88-8.08ZM160 121.63a6 6 0 0 0 6-6V72a6 6 0 0 0-12 0v43.63a6 6 0 0 0 6 6Zm40 44a6 6 0 0 0 6-6V32a6 6 0 0 0-12 0v127.63a6 6 0 0 0 6 6Z"/>`),
		g.Group(children),
	)
}