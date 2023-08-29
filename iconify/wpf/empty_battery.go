package wpf

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EmptyBattery(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<path fill="currentColor" d="M23 7c.551 0 1 .448 1 1v10c0 .552-.449 1-1 1H5c-.551 0-1-.448-1-1v-3H2v-4h2V8c0-.552.449-1 1-1h18m0-2H5a3 3 0 0 0-3 3v1H1c-.551 0-1 .449-1 1v6c0 .551.449 1 1 1h1v1a3 3 0 0 0 3 3h18a3 3 0 0 0 3-3V8a3 3 0 0 0-3-3z"/>`),
		g.Group(children),
	)
}