package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TaxiFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 112h-11.36l-27.39-47.94A16 16 0 0 0 187.36 56h-21.94l-12-29.94A15.93 15.93 0 0 0 138.58 16h-21.16a15.93 15.93 0 0 0-14.86 10.06L90.58 56H68.64a16 16 0 0 0-13.89 8.06L27.36 112H16a8 8 0 0 0 0 16h8v80a16 16 0 0 0 16 16h24a16 16 0 0 0 16-16v-16h96v16a16 16 0 0 0 16 16h24a16 16 0 0 0 16-16v-80h8a8 8 0 0 0 0-16ZM117.42 32h21.16l9.6 24h-40.36ZM80 160H64a8 8 0 0 1 0-16h16a8 8 0 0 1 0 16Zm112 0h-16a8 8 0 0 1 0-16h16a8 8 0 0 1 0 16Z"/>`),
		g.Group(children),
	)
}