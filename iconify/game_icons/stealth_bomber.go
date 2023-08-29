package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StealthBomber(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 32L20 400l60 64l52.1-75.9L176 432l50.5-50.5L256 448l29.5-66.5L336 432l43.9-43.9L432 464l60-64L256 32zm-9 47v78l-39-13l39-65zm18 0l39 65l-39 13V79z"/>`),
		g.Group(children),
	)
}