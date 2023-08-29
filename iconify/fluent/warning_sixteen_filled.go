package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WarningSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M9.092 2.638a1.25 1.25 0 0 0-2.182 0L2.157 11.14A1.25 1.25 0 0 0 3.247 13h9.504a1.25 1.25 0 0 0 1.091-1.86l-4.75-8.502ZM8.75 10.25a.75.75 0 1 1-1.5 0a.75.75 0 0 1 1.5 0ZM7.5 8V5.5a.5.5 0 0 1 1 0V8a.5.5 0 0 1-1 0Z"/>`),
		g.Group(children),
	)
}