package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Libby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m39.85 14.43l-15.42 3.64a1.79 1.79 0 0 1-.86 0L8.15 14.43a1.87 1.87 0 0 0-1.87 1.86V38a1.87 1.87 0 0 0 1.87 1.86l15.42 3.62a1.79 1.79 0 0 0 .86 0l15.42-3.62A1.87 1.87 0 0 0 41.72 38V16.29a1.87 1.87 0 0 0-1.87-1.86Zm-17.28 3.41v25.37m2.86-25.37v25.37"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.56 15a15.1 15.1 0 0 0-27.12 0"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.17 7.82V4.5l2.53 1.63l2.53-1.63v7.02"/>`),
		g.Group(children),
	)
}