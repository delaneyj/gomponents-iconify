package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LeftArrowCurvingRight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="m47 46.3l-9.1-9.7v6.1h-7.3c-5.3 0-9.7-5.1-9.7-11.2c0-6.2 4.8-11.3 10.8-11.3h9.2V12h-9.2C21.3 12 13 20.7 13 31.4c0 10.5 8 19.4 17.5 19.4h7.3V56l9.2-9.7z"/>`),
		g.Group(children),
	)
}