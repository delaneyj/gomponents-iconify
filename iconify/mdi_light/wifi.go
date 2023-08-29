package mdi_light

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wifi(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m11.5 20.5l-4.5-6c1.24-.94 2.81-1.5 4.5-1.5s3.26.56 4.5 1.5l-4.5 6m0-1.66l3.07-4.07a6.5 6.5 0 0 0-6.14 0l3.07 4.07M11.5 4c3.73 0 7.17 1.24 9.93 3.32l-.6.8C18.23 6.16 15 5 11.5 5S4.77 6.16 2.17 8.12l-.6-.8C4.33 5.24 7.77 4 11.5 4m0 6c2.37 0 4.56.79 6.32 2.11l-.6.8A9.458 9.458 0 0 0 11.5 11c-2.15 0-4.13.71-5.72 1.91l-.6-.8C6.94 10.79 9.13 10 11.5 10m0-3c3.05 0 5.86 1 8.12 2.72l-.62.78A12.45 12.45 0 0 0 11.5 8c-2.82 0-5.43.94-7.5 2.5l-.62-.78C5.64 8 8.45 7 11.5 7Z"/>`),
		g.Group(children),
	)
}