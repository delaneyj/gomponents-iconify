package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bloom(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBloom0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="m6 32l12 10m24-10L30 42m-6-10v12"/><path fill="#555" d="m17 11l7-7l7 7l7-1s1 4.239 1 7c0 10-8.5 15-15 15S9 27 9 17c0-2.761 1-7 1-7l7 1Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBloom0)"/>`),
		g.Group(children),
	)
}