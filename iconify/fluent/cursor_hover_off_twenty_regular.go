package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CursorHoverOffTwentyRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<path fill="currentColor" d="M2.854 2.146a.5.5 0 1 0-.708.708l1.242 1.241A2 2 0 0 0 2 6v7a2 2 0 0 0 2 2h5v-1H4a1 1 0 0 1-1-1V6a1 1 0 0 1 1-1h.293L10 10.707V17.5a.5.5 0 0 0 .91.287l1.571-2.245l2.781.427l1.884 1.885a.5.5 0 0 0 .708-.708l-15-15Zm11.212 12.628l-1.74-.268a.5.5 0 0 0-.486.207l-.84 1.2v-4.206l3.066 3.067Zm2.84.01A2 2 0 0 0 18 13V6a2 2 0 0 0-2-2H6.121l1 1H16a1 1 0 0 1 1 1v7a1 1 0 0 1-.885.993l.79.79Z"/>`),
		g.Group(children),
	)
}