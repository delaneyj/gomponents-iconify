package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cocktail(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M373 73q18-25 5-49q-11-21-37-21H41Q19 3 4 24Q-7 52 9 71l162 175v119q0 4-5 9t-20.5 9t-38.5 4q-22 0-22 21t22 21h170q22 0 22-21t-22-21q-23 0-38.5-4t-20.5-9t-5-9V246zM192 205L83 88h220zM341 45l-21 24v-2H64L45 45h296z"/>`),
		g.Group(children),
	)
}