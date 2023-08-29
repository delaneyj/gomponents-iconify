package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HashLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M224 90h-51l8.89-48.93a6 6 0 1 0-11.8-2.14L160.81 90H109l8.89-48.93a6 6 0 0 0-11.8-2.14L96.81 90H48a6 6 0 0 0 0 12h46.63l-9.46 52H32a6 6 0 0 0 0 12h51l-8.9 48.93a6 6 0 0 0 4.83 7A5.64 5.64 0 0 0 80 222a6 6 0 0 0 5.89-4.93l9.3-51.07H147l-8.89 48.93a6 6 0 0 0 4.83 7a5.64 5.64 0 0 0 1.08.1a6 6 0 0 0 5.89-4.93l9.28-51.1H208a6 6 0 0 0 0-12h-46.63l9.46-52H224a6 6 0 0 0 0-12Zm-74.83 64h-51.8l9.46-52h51.8Z"/>`),
		g.Group(children),
	)
}