package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WineglassLineDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><path fill="currentColor" d="M15.75 21.75a.75.75 0 0 0 0-1.5v1.5Zm-7.5-1.5a.75.75 0 0 0 0 1.5v-1.5Zm7.5 0H12v1.5h3.75v-1.5Zm-3.75 0H8.25v1.5H12v-1.5ZM6.895 3.75h10.21v-1.5H6.895v1.5ZM18.25 4.895V8h1.5V4.895h-1.5ZM5.75 8V4.895h-1.5V8h1.5ZM12 14.25A6.25 6.25 0 0 1 5.75 8h-1.5A7.75 7.75 0 0 0 12 15.75v-1.5ZM18.25 8A6.25 6.25 0 0 1 12 14.25v1.5A7.75 7.75 0 0 0 19.75 8h-1.5Zm-1.145-4.25c.633 0 1.145.513 1.145 1.145h1.5a2.645 2.645 0 0 0-2.645-2.645v1.5Zm-10.21-1.5A2.645 2.645 0 0 0 4.25 4.895h1.5c0-.632.513-1.145 1.145-1.145v-1.5Z"/><path fill="currentColor" d="M12.75 15.286a.75.75 0 0 0-1.5 0h1.5ZM11.25 21a.75.75 0 0 0 1.5 0h-1.5Zm0-5.714V21h1.5v-5.714h-1.5Z" opacity=".5"/><path stroke="currentColor" stroke-linecap="round" stroke-width="1.5" d="M5.5 9S7.581 8.087 9 8c2.465-.152 3.535 2.152 6 2c1.419-.087 3.5-1 3.5-1" opacity=".5"/></g>`),
		g.Group(children),
	)
}