package fa_6_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mound(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M144.1 179.2C173.8 127.7 228.6 96 288 96s114.2 31.7 143.9 83.2L540.4 368c12.3 21.3-3.1 48-27.7 48H63.3c-24.6 0-40-26.6-27.7-48l108.5-188.8z"/>`),
		g.Group(children),
	)
}