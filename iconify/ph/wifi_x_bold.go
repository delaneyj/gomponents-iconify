package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WifiXBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M144 204a16 16 0 1 1-16-16a16 16 0 0 1 16 16Zm73-124l15.52-15.51a12 12 0 0 0-17-17L200 63l-15.51-15.49a12 12 0 0 0-17 17L183 80l-15.49 15.51a12 12 0 0 0 17 17L200 97l15.51 15.52a12 12 0 0 0 17-17Zm-41.9 75.3a80 80 0 0 0-94.13 0a12 12 0 1 0 14.13 19.4a56 56 0 0 1 65.87 0a12 12 0 0 0 14.13-19.4ZM131.71 68h.3a12 12 0 0 0 .28-24H128A176.27 176.27 0 0 0 16.39 83.91a12 12 0 1 0 15.23 18.55A152.24 152.24 0 0 1 128 68h3.71Zm-.12 48a12 12 0 0 0 .82-24H128a126.66 126.66 0 0 0-79.45 27.64a12 12 0 0 0 14.9 18.81A102.89 102.89 0 0 1 128 116c1.2 0 2.41 0 3.59.06Z"/>`),
		g.Group(children),
	)
}