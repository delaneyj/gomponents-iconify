package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RightArrowCurvingLeft(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="m17 46.3l9.1-9.7v6.1h7.3c5.3 0 9.7-5.1 9.7-11.2c0-6.2-4.8-11.3-10.8-11.3h-9.2V12h9.2C42.7 12 51 20.7 51 31.4c0 10.5-8 19.4-17.5 19.4h-7.3V56L17 46.3z"/>`),
		g.Group(children),
	)
}