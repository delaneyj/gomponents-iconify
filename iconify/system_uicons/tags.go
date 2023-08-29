package system_uicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tags(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 21 21"),
		g.Raw(`<g fill="none" fill-rule="evenodd" transform="translate(1 3)"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.914.5H15.5a2 2 0 0 1 2 2v3.586a1 1 0 0 1-.293.707l-6.793 6.793a2 2 0 0 1-2.828 0l-3.172-3.172a2 2 0 0 1 0-2.828L11.207.793A1 1 0 0 1 11.914.5z"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.5 13.5l-2.013 1.006A2 2 0 0 1 2.72 13.42L1.105 9.114a2 2 0 0 1 .901-2.45L9.5 2.5"/><rect width="2" height="2" x="14" y="2" fill="currentColor" rx="1"/></g>`),
		g.Group(children),
	)
}