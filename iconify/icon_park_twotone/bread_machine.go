package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BreadMachine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTBreadMachine0"><g fill="none" stroke="#fff" stroke-width="4"><path fill="#555" stroke-linecap="round" stroke-linejoin="round" d="M7 27a8 8 0 0 1 8-8h18a8 8 0 0 1 8 8v8a4 4 0 0 1-4 4H11a4 4 0 0 1-4-4v-8Z"/><path stroke-linecap="round" stroke-linejoin="round" d="M34 19v-7c5 0 5-7 0-7H14c-5 0-5 7 0 7v7m8-8l-2 2m8-2l-2 2"/><circle cx="24" cy="29" r="4" fill="#555"/><path stroke-linecap="round" stroke-linejoin="round" d="M15 39v4m18-4v4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTBreadMachine0)"/>`),
		g.Group(children),
	)
}