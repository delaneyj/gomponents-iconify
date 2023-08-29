package et

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clipboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="currentColor"><path d="M1.5 32h25c.869 0 1.5-.631 1.5-1.5v-28c0-.869-.631-1.5-1.5-1.5h-4a.5.5 0 0 0 0 1h4c.313 0 .5.187.5.5v28c0 .313-.187.5-.5.5h-25c-.313 0-.5-.187-.5-.5v-28c0-.313.187-.5.5-.5h4a.5.5 0 0 0 0-1h-4C.631 1 0 1.631 0 2.5v28c0 .869.631 1.5 1.5 1.5z"/><path d="M5.5 5a.5.5 0 0 0 0-1h-2a.5.5 0 0 0-.5.5v24a.5.5 0 0 0 .5.5h21a.5.5 0 0 0 .5-.5v-24a.5.5 0 0 0-.5-.5h-2a.5.5 0 0 0 0 1H24v23H4V5h1.5z"/><path d="M9 7h10c1.215 0 2-.785 2-2V.5a.5.5 0 0 0-.5-.5h-13a.5.5 0 0 0-.5.5V5c0 1.215.785 2 2 2zM8 1h12v4c0 .664-.337 1-1 1H9c-.663 0-1-.336-1-1V1zm.5 15h11a.5.5 0 0 0 0-1h-11a.5.5 0 0 0 0 1zm0-4h11a.5.5 0 0 0 0-1h-11a.5.5 0 0 0 0 1zm0 8h11a.5.5 0 0 0 0-1h-11a.5.5 0 0 0 0 1zm0 4h11a.5.5 0 0 0 0-1h-11a.5.5 0 0 0 0 1z"/></g>`),
		g.Group(children),
	)
}