package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ArrowsClockwiseDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<g fill="currentColor"><path d="M216 48v48h-48ZM40 208l48-48H40Z" opacity=".2"/><path d="M219.06 40.61a8 8 0 0 0-8.72 1.73l-16.06 16.07C174.13 41.1 151.36 32 128 32c-42.82 0-68.58 25.27-69.66 26.34a8 8 0 0 0 11.32 11.32C69.87 69.44 91.73 48 128 48c23.17 0 41.92 10.85 54.92 21.76l-20.58 20.58A8 8 0 0 0 168 104h48a8 8 0 0 0 8-8V48a8 8 0 0 0-4.94-7.39ZM208 88h-20.69L208 67.31Zm-21.66 98.34c-.21.22-22.07 21.66-58.34 21.66c-23.17 0-41.92-10.85-54.92-21.76l20.58-20.58A8 8 0 0 0 88 152H40a8 8 0 0 0-8 8v48a8 8 0 0 0 13.66 5.66l16.06-16.07C81.87 214.9 104.64 224 128 224c42.82 0 68.58-25.27 69.66-26.34a8 8 0 0 0-11.32-11.32ZM48 168h20.69L48 188.69Z"/></g>`),
		g.Group(children),
	)
}