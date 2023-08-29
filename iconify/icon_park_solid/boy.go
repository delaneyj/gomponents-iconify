package icon_park_solid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Boy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipSBoy0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><circle cx="24" cy="10" r="6" fill="#fff"/><path fill="#fff" d="M30 36H18l-8-20h28l-8 20Z"/><path d="M27 36v8m-6-8v8"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipSBoy0)"/>`),
		g.Group(children),
	)
}