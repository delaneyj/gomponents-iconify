package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FlipVerticalSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M12.763 1.075A.5.5 0 0 1 13 1.5v5a.5.5 0 0 1-.5.5h-10a.5.5 0 0 1-.224-.947l10-5a.5.5 0 0 1 .487.022ZM4.618 6H12V2.309L4.618 6ZM13 14.5a.5.5 0 0 1-.724.447l-10-5A.5.5 0 0 1 2.5 9h10a.5.5 0 0 1 .5.5v5Z"/>`),
		g.Group(children),
	)
}