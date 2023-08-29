package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sparkasse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="30.1" height="26.04" x="8.95" y="17.46" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="5"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.97 26.14h23.08m-30.1 8.68h23.08"/><circle cx="23.98" cy="9.5" r="5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}