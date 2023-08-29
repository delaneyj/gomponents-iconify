package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StepsSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M5 13H2c-.552 0-1-.449-1-1V9c0-.551.448-1 1-1h3V6c0-.551.448-1 1-1h3V2c0-.551.448-1 1-1h3c.552 0 1 .449 1 1v4h-1V2h-3v3c0 .551-.448 1-1 1H6v2c0 .551-.448 1-1 1H2v3h3v1Zm10-.5V8c0-.551-.448-1-1-1h-3c-.552 0-1 .449-1 1v2H7c-.552 0-1 .449-1 1v3c0 .551.448 1 1 1h5.5c1.379 0 2.5-1.122 2.5-2.5Z"/>`),
		g.Group(children),
	)
}