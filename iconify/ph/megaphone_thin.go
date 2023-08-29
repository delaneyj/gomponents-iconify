package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MegaphoneThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M192 76h-39.9c-2.19-.11-54.55-3.38-104.38-45.17A12 12 0 0 0 28 40v160a11.88 11.88 0 0 0 6.92 10.87A12.18 12.18 0 0 0 40 212a11.93 11.93 0 0 0 7.69-2.83c42.46-35.62 86.76-43.25 100.28-44.8v36.29a12 12 0 0 0 5.34 10l11 7.33a12 12 0 0 0 18.28-7l12.49-47.09A44 44 0 0 0 192 76Zm-44 80.3c-13.53 1.35-60 8.62-105.42 46.75A4 4 0 0 1 36 200V40a4 4 0 0 1 2.3-3.62a4.32 4.32 0 0 1 1.75-.38a3.88 3.88 0 0 1 2.53 1C88 75.09 134.47 82.36 148 83.71Zm26.87 52.7a4 4 0 0 1-6.1 2.36l-11-7.33a4 4 0 0 1-1.78-3.33V164h30.8ZM192 156h-36V84h36a36 36 0 0 1 0 72Z"/>`),
		g.Group(children),
	)
}