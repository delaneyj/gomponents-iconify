package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Figma(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke-linejoin="round" stroke-width="4"><rect width="36" height="36" x="6" y="6" fill="#2F88FF" stroke="#000" rx="3"/><path fill="#43CCF8" stroke="#fff" d="M18 18C18 16.3431 19.3431 15 21 15H24V21H21C19.3431 21 18 19.6569 18 18Z"/><path fill="#43CCF8" stroke="#fff" d="M18 24C18 22.3431 19.3431 21 21 21H24V27H21C19.3431 27 18 25.6569 18 24Z"/><path fill="#43CCF8" stroke="#fff" d="M18 30C18 28.3431 19.3431 27 21 27H24V30C24 31.6569 22.6569 33 21 33C19.3431 33 18 31.6569 18 30Z"/><path fill="#43CCF8" stroke="#fff" d="M24 15H27C28.6569 15 30 16.3431 30 18C30 19.6569 28.6569 21 27 21H24V15Z"/><path fill="#43CCF8" stroke="#fff" d="M24 24C24 22.3431 25.3431 21 27 21C28.6569 21 30 22.3431 30 24C30 25.6569 28.6569 27 27 27C25.3431 27 24 25.6569 24 24Z"/></g>`),
		g.Group(children),
	)
}