package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonAvailableTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M17.5 12a5.5 5.5 0 1 1 0 11a5.5 5.5 0 0 1 0-11Zm-5.477 2A6.47 6.47 0 0 0 11 17.5c0 1.63.6 3.12 1.592 4.262c-.795.16-1.66.24-2.592.24c-3.42 0-5.944-1.073-7.486-3.237a2.75 2.75 0 0 1-.51-1.596v-.92A2.249 2.249 0 0 1 4.253 14h7.77Zm2.83 3.147a.5.5 0 1 0-.706.707l2 2a.5.5 0 0 0 .707 0l4-4a.5.5 0 1 0-.707-.707L16.5 18.793l-1.646-1.646ZM10 2.005a5 5 0 1 1 0 10a5 5 0 0 1 0-10Z"/>`),
		g.Group(children),
	)
}