package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gym(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="m8.11 15.89l7.779-7.78M2 22l1.848-1.848M22 2l-1.848 1.848M2.1 9.878L14.121 21.9l-.177.177l-1.737-.208c-1.75-.21-3.522.138-5.098.926L6.697 23L1 17.303l.206-.412c.788-1.576 1.136-3.348.926-5.098l-.208-1.737l.176-.177Zm7.78-7.776l12.02 12.02l.177-.177l-.208-1.737c-.21-1.75.138-3.522.926-5.098L23 6.697L17.303 1l-.413.206c-1.575.788-3.348 1.136-5.097.926l-1.738-.208l-.176.177Z"/>`),
		g.Group(children),
	)
}