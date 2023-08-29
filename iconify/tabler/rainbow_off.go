package tabler

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RainbowOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M22 17c0-5.523-4.477-10-10-10c-.308 0-.613.014-.914.041m-3.208.845A10 10 0 0 0 2 17m9.088-5.931A6 6 0 0 0 6 17m8 0a2 2 0 1 0-4 0M3 3l18 18"/>`),
		g.Group(children),
	)
}