package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTDrink0"><g fill="none" stroke="#fff" stroke-width="4"><path stroke-linecap="round" d="M10 16h28"/><path d="M14.153 18.142A2 2 0 0 1 16.148 16h15.704a2 2 0 0 1 1.995 2.142l-1.714 24A2 2 0 0 1 30.138 44H17.862a2 2 0 0 1-1.995-1.858l-1.714-24Z"/><path stroke-linecap="round" d="M24 10V6a2 2 0 0 1 2-2h3"/><path fill="#555" d="M14.722 11.671A2 2 0 0 1 16.694 10h14.612a2 2 0 0 1 1.973 1.671L34 16H14l.722-4.329Z"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDrink0)"/>`),
		g.Group(children),
	)
}