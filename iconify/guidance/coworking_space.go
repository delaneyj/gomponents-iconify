package guidance

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CoworkingSpace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" d="M6.5 21a7.526 7.526 0 0 0-.847-3.469A3.227 3.227 0 0 0 5.3 17H.5v.25c.655 1.147 1 2.43 1 3.75H10m5 2.5H0m15.5-14a7.527 7.527 0 0 0-.847-3.469A3.224 3.224 0 0 0 14.3 5.5H9.5v.25c.655 1.147 1 2.43 1 3.75H19m5 2.5h-9m-.5 9.5V19s-1.5-1-4-1c-1.006 0-1.85.162-2.5.355M23.5 10V7.5s-1.5-1-4-1c-1.006 0-1.85.162-2.5.356M10.35 16s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0C12.246 15 10.65 16 10.65 16h-.3Zm9-11.5s-1.6-1-1.6-2.25a1.747 1.747 0 1 1 3.496 0c0 1.25-1.596 2.25-1.596 2.25h-.3Z"/>`),
		g.Group(children),
	)
}