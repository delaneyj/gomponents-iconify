package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipVerticalSixteenFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.651 1.116a.75.75 0 0 1 .35.634v4.5a.75.75 0 0 1-.75.75h-9.5a.75.75 0 0 1-.322-1.428l9.5-4.5a.75.75 0 0 1 .722.044ZM6.085 5.5H11.5V2.935L6.085 5.5Zm6.915 9a.5.5 0 0 1-.724.447l-10-5A.5.5 0 0 1 2.5 9h10a.5.5 0 0 1 .5.5v5Z"/>`),
		g.Group(children),
	)
}