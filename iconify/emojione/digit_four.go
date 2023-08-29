package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DigitFour(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#4fd1d9"/><path fill="#fff" d="M33.7 48v-6.4H20v-5.3L34.5 16h5.4v20.2H44v5.4h-4.1V48h-6.2zm0-11.8V25.3L26 36.2h7.7z"/>`),
		g.Group(children),
	)
}