package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Entertainment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTEntertainment0"><g fill="none" stroke="#fff" stroke-linejoin="round" stroke-width="4"><path fill="#555" d="M16 24c5.523 0 10-4.477 10-10S21.523 4 16 4S6 8.477 6 14s4.477 10 10 10Z"/><path stroke-linecap="round" d="M26 15.202c.014.014 4.723 5.936 14.126 17.764a1 1 0 0 1-.062 1.346l-4.084 4.084a1 1 0 0 1-1.346.062L17.822 24m8.648.47l2.829 2.829"/><path stroke-linecap="round" d="M17 44.086c1.917-2.498 4.247-3.747 6.99-3.747c4.116 0 8.973 5.17 13.116 4.459c4.144-.712 5.33-4.798 2.78-7.06"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTEntertainment0)"/>`),
		g.Group(children),
	)
}