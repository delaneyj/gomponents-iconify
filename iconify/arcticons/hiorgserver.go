package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Hiorgserver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.907 22.576H17.093a.837.837 0 0 0-.838.837h0v13.815c0 .462.375.837.838.837h13.814a.837.837 0 0 0 .837-.837V23.413a.837.837 0 0 0-.837-.837h0Z"/><circle cx="7.066" cy="28.307" r="2.566" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="40.934" cy="28.307" r="2.566" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="24" cy="12.502" r="2.566" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="11.172" cy="16.74" r="2.566" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.828" cy="16.74" r="2.566" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}