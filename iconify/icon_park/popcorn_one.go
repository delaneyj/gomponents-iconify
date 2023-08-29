package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PopcornOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" stroke="#000" d="M7 16H41L34 44H14L7 16Z"/><path stroke="#fff" d="M20 16V44"/><path stroke="#fff" d="M28 16V44"/><path fill="#2F88FF" stroke="#000" d="M33 9C30.7909 9 29 10.7909 29 13C29 14.1947 29.5238 15.2671 30.3542 16H35.6458C36.4762 15.2671 37 14.1947 37 13C37 10.7909 35.2091 9 33 9Z"/><path fill="#2F88FF" stroke="#000" d="M24 9C21.7909 9 20 10.7909 20 13C20 14.1947 20.5238 15.2671 21.3542 16H26.6458C27.4762 15.2671 28 14.1947 28 13C28 10.7909 26.2091 9 24 9Z"/><path fill="#2F88FF" stroke="#000" d="M15 9C12.7909 9 11 10.7909 11 13C11 14.1947 11.5238 15.2671 12.3542 16H17.6458C18.4762 15.2671 19 14.1947 19 13C19 10.7909 17.2091 9 15 9Z"/><path stroke="#000" d="M22.874 9C22.9562 8.68038 23 8.3453 23 8C23 5.79086 21.2091 4 19 4C16.7909 4 15 5.79086 15 8C15 8.3453 15.0438 8.68038 15.126 9"/><path stroke="#000" d="M32.874 9C32.9562 8.68038 33 8.3453 33 8C33 5.79086 31.2091 4 29 4C26.7909 4 25 5.79086 25 8C25 8.3453 25.0438 8.68038 25.126 9"/><path stroke="#000" d="M16 16L32 16"/><path stroke="#000" d="M16 44L32 44"/></g>`),
		g.Group(children),
	)
}