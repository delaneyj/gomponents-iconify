package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Howl(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M2.73 25.253s4.693 1.177 6.52 0c5.522-3.555.73-17.554 8.153-17.935c12.621-.533-1.573 34.217 12.487 34.239c7.042-.382 3.54-13.11 8.709-16.304c1.85-1.143 6.522 0 6.522 0"/>`),
		g.Group(children),
	)
}