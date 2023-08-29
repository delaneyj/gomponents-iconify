package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SyncOffSixteenRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M14 8A6 6 0 1 1 2 8a6 6 0 0 1 12 0Zm-1 0H9.941l-.971 2.67a.5.5 0 1 1-.94-.34l2-5.5a.5.5 0 0 1 .94.34L10.305 7H12.9a5.002 5.002 0 0 0-9.8 0h3.14l.79-2.17a.5.5 0 0 1 .94.34l-2 5.5a.5.5 0 1 1-.94-.34L5.877 8H3a5 5 0 0 0 10 0Z"/>`),
		g.Group(children),
	)
}