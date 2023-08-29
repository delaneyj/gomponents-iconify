package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Congstar(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m33.056 14.4l9.705-3.668c-4.566-5.165-12.303-8.747-19.243-8.17a24.007 24.007 0 0 0-18.08 11.504a21.572 21.572 0 0 0-.25 19.426a23.631 23.631 0 0 0 17.998 11.923c7.06.541 11.41-1.556 16.174-4.503A16.194 16.194 0 0 0 45 34.326l-8.543-2.668c-.812-.254-1.786-.503-2.405.75l-4.904 9.932a1.262 1.262 0 0 1-2.385-.25l-2.27-12.006a2.265 2.265 0 0 0-1.887-2.178L9.833 24.655c-.921-.235-.64-1.536 0-1.751l11.02-3.72c1.009-.34 1.144-1.213.996-2.116l-1.814-11.09c-.132-.802 1.153-1.54 1.731-.937l8.553 8.942a2.616 2.616 0 0 0 2.737.417Z"/>`),
		g.Group(children),
	)
}