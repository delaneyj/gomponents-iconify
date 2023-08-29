package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M21.485 16.155a.992.992 0 0 0-.77-.36h-.33v-9.23a2.5 2.5 0 0 0-2.5-2.5H6.115a2.5 2.5 0 0 0-2.5 2.5V15.8h-.34a1 1 0 0 0-.98 1.17l.3 1.73a1.5 1.5 0 0 0 1.48 1.24h15.85a1.5 1.5 0 0 0 1.48-1.24l.3-1.73a.986.986 0 0 0-.22-.815Zm-16.87-9.59a1.5 1.5 0 0 1 1.5-1.5h11.77a1.5 1.5 0 0 1 1.5 1.5V15.8H4.615Zm15.8 11.96a.494.494 0 0 1-.49.41H4.075a.494.494 0 0 1-.49-.41l-.31-1.73h17.44Z"/>`),
		g.Group(children),
	)
}