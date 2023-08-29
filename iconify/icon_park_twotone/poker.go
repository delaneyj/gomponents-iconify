package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Poker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTPoker0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M42 4H12v40h30V4Z"/><path stroke-linecap="round" d="M4 11.79L12 10v34L4 11.79Z" clip-rule="evenodd"/><path fill="#555" d="m27 18l-5 6l5 6l5-6l-5-6Z"/><path stroke-linecap="round" d="M18 10v4m18 20v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTPoker0)"/>`),
		g.Group(children),
	)
}