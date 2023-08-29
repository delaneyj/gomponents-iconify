package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShirtFoldedDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M208 64v152a8 8 0 0 1-8 8H56a8 8 0 0 1-8-8V64a8 8 0 0 1 8-8h24v56a8 8 0 0 0 13.12 6.15L128 88l34.88 30.13A8 8 0 0 0 176 112V56h24a8 8 0 0 1 8 8ZM96 40l32 48l32-48Z" opacity=".2"/><path d="M200 48h-20.69l-13.65-13.66A8.08 8.08 0 0 0 160 32H95.28a8 8 0 0 0-4.89 2.29L76.69 48H56a16 16 0 0 0-16 16v152a16 16 0 0 0 16 16h144a16 16 0 0 0 16-16V64a16 16 0 0 0-16-16Zm-32 11.31V112l-29.43-25.44l22.67-34Zm-40 14.27L111 48h34.1Zm-33.24-21l22.67 34L88 112V59.31ZM56 64h16v48a15.85 15.85 0 0 0 9.21 14.49A16.1 16.1 0 0 0 88 128a15.89 15.89 0 0 0 10.2-3.73a.52.52 0 0 0 .11-.1L120 105.48V216H56Zm144 152h-64V105.48l21.65 18.7a.52.52 0 0 0 .11.1A15.89 15.89 0 0 0 168 128a16.1 16.1 0 0 0 6.83-1.54A15.85 15.85 0 0 0 184 112V64h16Z"/></g>`),
		g.Group(children),
	)
}