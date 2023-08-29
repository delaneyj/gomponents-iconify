package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Volumeup(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7.5 11.5c-.33 0-.65-.11-.92-.32L3.83 9.04H2.5c-.83 0-1.5-.67-1.5-1.5v-2c0-.83.67-1.5 1.5-1.5h1.33L6.58 1.9c.46-.35 1.06-.42 1.58-.16c.52.25.84.77.84 1.35V10a1.495 1.495 0 0 1-1.5 1.5Zm-5-6.46c-.28 0-.5.22-.5.5v2c0 .28.22.5.5.5H4c.11 0 .22.04.31.11l2.89 2.24c.15.12.35.14.53.05c.18-.09.28-.25.28-.45v-6.9c0-.2-.1-.36-.28-.45a.487.487 0 0 0-.53.05L4.31 4.93a.51.51 0 0 1-.31.11H2.5Zm10 4c-.28 0-.5-.22-.5-.5v-4c0-.28.22-.5.5-.5s.5.22.5.5v4c0 .28-.22.5-.5.5Z"/><path fill="currentColor" d="M14.5 7.04h-4c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h4c.28 0 .5.22.5.5s-.22.5-.5.5Z"/>`),
		g.Group(children),
	)
}