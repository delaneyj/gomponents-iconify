package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Whirlwind(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTWhirlwind0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M24 29a5 5 0 1 0 0-10a5 5 0 0 0 0 10Z"/><path d="M23.5 44C16.596 44 11 38.404 11 31.5S16.596 19 23.5 19"/><path d="M44 23.5C44 30.404 38.404 36 31.5 36S19 30.404 19 23.5"/><path d="M23.5 29C30.404 29 36 23.404 36 16.5S30.404 4 23.5 4"/><path d="M29 23.5C29 16.596 23.404 11 16.5 11S4 16.596 4 23.5"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTWhirlwind0)"/>`),
		g.Group(children),
	)
}