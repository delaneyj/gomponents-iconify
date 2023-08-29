package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DnaLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M198 204.5V232a6 6 0 0 1-12 0v-27.5a65.64 65.64 0 0 0-36.48-59l-48.4-24.2A77.57 77.57 0 0 1 58 51.5V24a6 6 0 0 1 12 0v27.5a65.64 65.64 0 0 0 36.48 59l48.4 24.2A77.57 77.57 0 0 1 198 204.5Zm-38-2.5H70.05A66 66 0 0 1 74 182h74.13a6 6 0 0 0 0-12H79.77a65.85 65.85 0 0 1 17.16-18.7a6 6 0 0 0-7.1-9.67A78.27 78.27 0 0 0 58 204.5V232a6 6 0 0 0 12 0v-18h90a6 6 0 0 0 0-12Zm32-184a6 6 0 0 0-6 6v18H96a6 6 0 0 0 0 12h90a66 66 0 0 1-4 20h-74.11a6 6 0 1 0 0 12h68.34a65.85 65.85 0 0 1-17.16 18.7a6 6 0 0 0 7.1 9.67A78.27 78.27 0 0 0 198 51.5V24a6 6 0 0 0-6-6Z"/>`),
		g.Group(children),
	)
}