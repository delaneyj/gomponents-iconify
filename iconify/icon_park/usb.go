package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Usb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="4"><path fill="#2F88FF" d="M12 22C14.2091 22 16 20.2091 16 18C16 15.7909 14.2091 14 12 14C9.79086 14 8 15.7909 8 18C8 20.2091 9.79086 22 12 22Z"/><path fill="#2F88FF" d="M36 28C38.2091 28 40 26.2091 40 24C40 21.7909 38.2091 20 36 20C33.7909 20 32 21.7909 32 24C32 26.2091 33.7909 28 36 28Z"/><path stroke-linecap="round" d="M19 9L24 4L29 9"/><path stroke-linecap="round" d="M25 39L12 28.2632V22"/><path stroke-linecap="round" d="M36 28V32.7895L24 41"/><path stroke-linecap="round" d="M24 4V43"/><path stroke-linecap="round" d="M21 44H27"/></g>`),
		g.Group(children),
	)
}