package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Snappdriver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><circle cx="21.211" cy="36.036" r="4.251"/><path d="m31.435 7.759l-3.868 14.67c-.927 3.517-1.501 6.386-5.181 6.421l-3.598.034l5.66-21.119l6.987-.006Z"/></g>`),
		g.Group(children),
	)
}