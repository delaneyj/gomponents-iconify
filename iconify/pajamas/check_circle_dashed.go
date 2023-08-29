package pajamas

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CheckCircleDashed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M6.701 1.08a.75.75 0 0 0-.934-.5A7.75 7.75 0 0 0 .573 5.785a.75.75 0 1 0 1.438.429a6.25 6.25 0 0 1 4.188-4.199a.75.75 0 0 0 .502-.934Zm2.598 0a.75.75 0 0 1 .934-.501a7.75 7.75 0 0 1 5.194 5.206a.75.75 0 1 1-1.438.429a6.25 6.25 0 0 0-4.188-4.199a.75.75 0 0 1-.502-.934Zm2.481 4.14a.75.75 0 0 1 0 1.06l-4.5 4.5a.75.75 0 0 1-1.06 0L4.47 9.03a.75.75 0 0 1 1.06-1.06l1.22 1.22l3.97-3.97a.75.75 0 0 1 1.06 0Zm-1.547 10.2a.75.75 0 1 1-.432-1.436a6.251 6.251 0 0 0 4.188-4.199a.75.75 0 1 1 1.438.429a7.75 7.75 0 0 1-5.194 5.206Zm-4.466 0a.75.75 0 1 0 .432-1.436a6.25 6.25 0 0 1-4.188-4.199a.75.75 0 1 0-1.438.429a7.75 7.75 0 0 0 5.194 5.206Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}