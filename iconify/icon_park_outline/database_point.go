package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatabasePoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<defs><path id="ipODatabasePoint0" d="M39 9c0 2.761-6.716 5-15 5c-8.284 0-15-2.239-15-5s6.716-5 15-5c8.284 0 15 2.239 15 5Z"/><path id="ipODatabasePoint1" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M39 9c0 2.761-6.716 5-15 5c-8.284 0-15-2.239-15-5s6.716-5 15-5c8.284 0 15 2.239 15 5Z"/></defs><g fill="none"><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 36v-6m-4 10H6m22 0h14m-14 0a4 4 0 1 1-8 0a4 4 0 0 1 8 0ZM39 9v16c0 2.761-6.716 5-15 5c-8.284 0-15-2.239-15-5V9"/><path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M39 17c0 2.761-6.716 5-15 5c-8.284 0-15-2.239-15-5"/><use href="#ipODatabasePoint0"/><use href="#ipODatabasePoint0"/><use href="#ipODatabasePoint1" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><use href="#ipODatabasePoint1" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/></g>`),
		g.Group(children),
	)
}