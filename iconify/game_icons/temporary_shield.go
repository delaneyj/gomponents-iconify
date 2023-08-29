package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TemporaryShield(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 22C192 70 128 90 32 122c0 112 80 272 224 368c144-96 224-250 224-362c-96-32-160-58-224-106zm0 97c75.6 0 137 61.4 137 137s-61.4 137-137 137s-137-61.4-137-137s61.4-137 137-137zm-9 18.3c-42.9 3.2-79.3 29.1-97.6 65.6l97.9 39.9a16 17.12 0 0 1 8.7-2.8a16 17.12 0 0 1 2.9.3l38.7-38.7l12.8 12.8l-38.8 38.8a16 17.12 0 0 1 .4 3.9a16 17.12 0 0 1-16 17.1a16 17.12 0 0 1-15.9-15l-97.4-39.7c-2.8 8.8-4.7 18-5.4 27.5H160v18h-22.7c4.4 58.6 51.1 105.3 109.7 109.7V352h18v22.7c58.6-4.4 105.3-51.1 109.7-109.7H352v-18h22.7c-4.4-58.6-51.1-105.3-109.7-109.7V160h-18v-22.7z"/>`),
		g.Group(children),
	)
}