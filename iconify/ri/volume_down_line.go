package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VolumeDownLine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M13 7.22L9.603 10H6v4h3.603L13 16.78V7.22ZM8.889 16H5a1 1 0 0 1-1-1V9a1 1 0 0 1 1-1h3.889l5.294-4.332a.5.5 0 0 1 .817.387v15.89a.5.5 0 0 1-.817.387L8.89 16Zm9.974.591l-1.422-1.422A3.994 3.994 0 0 0 19 12c0-1.43-.75-2.685-1.88-3.392l1.439-1.439A5.991 5.991 0 0 1 21 12c0 1.842-.83 3.49-2.137 4.591Z"/>`),
		g.Group(children),
	)
}