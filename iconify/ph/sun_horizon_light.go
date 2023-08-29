package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SunHorizonLight(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M240 154h-42.72a70.91 70.91 0 0 0 .72-10a70 70 0 0 0-140 0a70.91 70.91 0 0 0 .72 10H16a6 6 0 0 0 0 12h224a6 6 0 0 0 0-12ZM70 144a58 58 0 1 1 115.13 10H70.87a58.63 58.63 0 0 1-.87-10Zm144 56a6 6 0 0 1-6 6H48a6 6 0 0 1 0-12h160a6 6 0 0 1 6 6ZM74.63 42.69a6 6 0 0 1 10.74-5.37l8 16a6 6 0 0 1-10.74 5.36Zm-56 50.63a6 6 0 0 1 8.05-2.69l16 8a6 6 0 0 1-5.36 10.74l-16-8a6 6 0 0 1-2.69-8.05Zm192 13.36a6 6 0 0 1 2.69-8.05l16-8a6 6 0 1 1 5.36 10.74l-16 8a6 6 0 0 1-8.05-2.69Zm-48-53.36l8-16a6 6 0 0 1 10.74 5.37l-8 16a6 6 0 1 1-10.74-5.36Z"/>`),
		g.Group(children),
	)
}