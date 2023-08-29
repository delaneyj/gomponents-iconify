package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bompass(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24.039" cy="23.76" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.611 10.373v8.131l-7.057 4.1l.017-8.163l7.04-4.068ZM15.389 34.977l4-6.928l.59 9.578l4-6.928"/>`),
		g.Group(children),
	)
}