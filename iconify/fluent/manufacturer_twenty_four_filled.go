package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ManufacturerTwentyFourFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M2.924 4.129A2.25 2.25 0 0 1 5.171 2H6.33a2.25 2.25 0 0 1 2.247 2.129l.838 15.5A2.25 2.25 0 0 1 7.167 22H4.333a2.25 2.25 0 0 1-2.246-2.371l.837-15.5z" fill="currentColor"/><path d="M9.512 22c.6-.625.95-1.488.9-2.425L9.856 9.284l4.666-3.862A.75.75 0 0 1 15.75 6v4.24l4.98-4.781A.75.75 0 0 1 22 6v13.75A2.25 2.25 0 0 1 19.75 22H9.512zM13 20.5h5V16a1 1 0 0 0-1-1h-3a1 1 0 0 0-1 1v4.5z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}