package game_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BusDoors(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M41 25v254h58.73l63.97 64H247V25H41zm224 0v318h83.3l64-64H471V25H265zM57 41h174v208H57V41zm224 0h174v208H281V41zM160 279h64v18h-64v-18zm128 0h64v18h-64v-18zM41 297v46h49.63l63.97 64h92.3v18h-99.7l-64.03-64H41v126h206V361h-90.7l-64.03-64H41zm378.7 0l-64 64H265v126h206V361h-42.2l-64 64h-99.7v-18h92.3l64-64H471v-46h-51.3z"/>`),
		g.Group(children),
	)
}