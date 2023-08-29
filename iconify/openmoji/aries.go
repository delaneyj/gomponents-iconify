package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Aries(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#b399c8" d="M12 12h48v48H12z"/><g fill="none" stroke="#000" stroke-linecap="round"><path stroke-miterlimit="10" stroke-width="3" d="M36 48.84c3.087-16.36 6.337-26.33 11.64-25.68c2.611.586 3.711 3.609 2.58 5.817"/><path stroke-linejoin="round" stroke-width="2" d="M12 12h48v48H12z"/><path stroke-miterlimit="10" stroke-width="3" d="M36 48.84c-3.087-16.36-6.337-26.33-11.64-25.68c-2.611.586-3.711 3.609-2.58 5.817"/></g>`),
		g.Group(children),
	)
}