package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Grenade(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 16c-96 0-120 48.188-120 72.188v71.718c0 13.296 10.61 23.907 23.907 23.907h23.907v144.375h-23.907c-13.296 0-23.907 10.61-23.907 23.907v71.717c0 24 24 72.188 120 72.188s120-48.188 120-72.188v-71.718c0-13.296-10.61-23.907-23.907-23.907h-23.907V183.813h23.907c13.296 0 23.907-10.61 23.907-23.907V88.187C376 64.188 352 16 256 16zm-23.907 167.813h47.813v144.375h-47.813V183.813z"/>`),
		g.Group(children),
	)
}