package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FireHazard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M12 23.5c-1.933 0-3.5-1.24-3.5-3.07c0-.762.32-1.588.89-2.127l1.455-1.378c.74-.7 1.155-1.93 1.155-2.92c0 .99.416 2.22 1.155 2.92l1.457 1.378c.569.54.889 1.365.889 2.127c0 1.83-1.567 3.07-3.5 3.07Zm0 0c-4.556 0-8.25-2.972-8.25-7.5c0-1.885.754-4.027 2.096-5.36l3.432-3.41C11.02 5.496 12 2.45 12 0c0 2.45.98 5.497 2.723 7.23l3.431 3.41c1.342 1.333 2.096 3.475 2.096 5.36c0 4.528-3.693 7.5-8.25 7.5Z"/>`),
		g.Group(children),
	)
}