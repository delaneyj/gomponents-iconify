package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatabasePoint(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<defs><path id="ipTDatabasePoint0" fill="#555" d="M39 9c0 2.761-6.716 5-15 5c-8.284 0-15-2.239-15-5s6.716-5 15-5c8.284 0 15 2.239 15 5Z"/><path id="ipTDatabasePoint1" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M39 9c0 2.761-6.716 5-15 5c-8.284 0-15-2.239-15-5s6.716-5 15-5c8.284 0 15 2.239 15 5Z"/></defs><mask id="ipTDatabasePoint2"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 36v-6m-4 10H6m22 0h14"/><path fill="#555" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M28 40a4 4 0 1 1-8 0a4 4 0 0 1 8 0Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M39 9v16c0 2.761-6.716 5-15 5c-8.284 0-15-2.239-15-5V9"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M39 17c0 2.761-6.716 5-15 5c-8.284 0-15-2.239-15-5"/><use href="#ipTDatabasePoint0"/><use href="#ipTDatabasePoint0"/><use href="#ipTDatabasePoint1" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/><use href="#ipTDatabasePoint1" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTDatabasePoint2)"/>`),
		g.Group(children),
	)
}