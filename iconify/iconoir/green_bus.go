package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GreenBus(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path stroke-linejoin="round" d="m7 16.01l.01-.011M21 12H3v7a1 1 0 0 0 1 1h9m4 3s.9-3.118 3-5"/><path stroke-linejoin="round" d="m19.802 21.424l-.134.013a3.094 3.094 0 0 1-3.366-2.774a3.06 3.06 0 0 1 2.761-3.35l2.986-.28a.35.35 0 0 1 .381.314l.255 2.58a3.194 3.194 0 0 1-2.883 3.497ZM21 8V6a4 4 0 0 0-4-4H7a4 4 0 0 0-4 4v2m4 0h10"/><path d="M4.5 20v1.9a.6.6 0 0 0 .6.6h2.8a.6.6 0 0 0 .6-.6V20"/></g>`),
		g.Group(children),
	)
}