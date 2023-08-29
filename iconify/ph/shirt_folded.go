package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShirtFolded(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 48h-20.69l-13.65-13.66a8.07 8.07 0 0 0-5.61-2.34H96a8 8 0 0 0-5.66 2.34L76.69 48H56a16 16 0 0 0-16 16v152a16 16 0 0 0 16 16h144a16 16 0 0 0 16-16V64a16 16 0 0 0-16-16Zm-38.76 4.56l6.76 6.75V112l-29.43-25.44ZM88 59.31l6.76-6.75l22.67 34L88 112ZM120 216H56V64h16v48a15.85 15.85 0 0 0 9.21 14.49A16.1 16.1 0 0 0 88 128a15.89 15.89 0 0 0 10.2-3.73a.52.52 0 0 0 .11-.1L120 105.48Zm-9-168h34.1L128 73.58Zm89 168h-64V105.48l21.65 18.7a.52.52 0 0 0 .11.1A15.89 15.89 0 0 0 168 128a16.1 16.1 0 0 0 6.83-1.54A15.85 15.85 0 0 0 184 112V64h16Z"/>`),
		g.Group(children),
	)
}