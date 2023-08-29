package zmdi

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Twitch(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 432 384"),
		g.Raw(`<path fill="currentColor" d="M296 374h-83l-56 55h-55v-55H0V77L28 3h379v259zm74-130V40H65v269h83v55l56-55h102zM269 114h37v111h-37V114zM167 225V114h37v111h-37z"/>`),
		g.Group(children),
	)
}