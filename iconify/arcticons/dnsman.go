package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dnsman(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.64 18.75h14.72m-14.72-3.76h14.72m-14.72-3.76h14.72M16.64 7.47h14.72M33.75 25a2 2 0 0 0-2-2h-15.5a2 2 0 0 0-2 2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42 25h-8.25V7.33a2.83 2.83 0 0 0-2.83-2.83H17.08a2.83 2.83 0 0 0-2.83 2.83V25H6a1.47 1.47 0 0 0 0 2.9h8.29v12.77a2.83 2.83 0 0 0 2.83 2.83h13.8a2.83 2.83 0 0 0 2.83-2.83V27.9H42a1.47 1.47 0 0 0 0-2.9Z"/>`),
		g.Group(children),
	)
}