package formkit

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Volumedown(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M7.5 11.46c-.33 0-.65-.11-.92-.32L3.83 9H2.5C1.67 9 1 8.33 1 7.5v-2C1 4.67 1.67 4 2.5 4h1.33l2.75-2.14c.46-.35 1.06-.42 1.58-.16c.52.25.84.77.84 1.35v6.91a1.495 1.495 0 0 1-1.5 1.5ZM2.5 5c-.28 0-.5.22-.5.5v2c0 .28.22.5.5.5H4c.11 0 .22.04.31.11l2.89 2.24c.15.12.35.14.53.05c.18-.09.28-.25.28-.45V3.04c0-.2-.1-.36-.28-.45a.487.487 0 0 0-.53.05L4.31 4.88a.51.51 0 0 1-.31.11H2.5Zm12 2h-4c-.28 0-.5-.22-.5-.5s.22-.5.5-.5h4c.28 0 .5.22.5.5s-.22.5-.5.5Z"/>`),
		g.Group(children),
	)
}