package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenNibStraightBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M225.92 122.11c-.1-.19-.19-.37-.3-.56L196 68.85V32a20 20 0 0 0-20-20H80a20 20 0 0 0-20 20v36.87l-29.62 52.68a5.12 5.12 0 0 0-.3.57a19.89 19.89 0 0 0 1.83 20.71c.09.13.18.25.28.37l86.44 108.29a12 12 0 0 0 18.75 0l86.43-108.29c.1-.12.19-.24.28-.37a19.91 19.91 0 0 0 1.83-20.72ZM172 36v24H84V36Zm-44 108a10 10 0 1 1 10-10a10 10 0 0 1-10 10Zm12 65.74v-43.93a34 34 0 1 0-24 0v43.92l-63.17-79.14L79 84h98l26.19 46.59Z"/>`),
		g.Group(children),
	)
}