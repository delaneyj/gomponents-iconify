package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Sagittarius(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#b399c8" d="M12 12h48v48H12z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round"><path stroke-width="2" d="M12 12h48v48H12z"/><path stroke-width="3" d="M37.81 24.19h9.996v9.996M24.19 47.81l23.62-23.62m-9.4 19.25l-9.99-9.98"/></g>`),
		g.Group(children),
	)
}