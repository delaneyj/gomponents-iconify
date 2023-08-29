package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mapy(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M5.88 42.05c5.91-13.85 9.24-22.24 7.82-33c-.77-5.81 3.94-5.36 6.09-2.68c2.71 3.38 6.63 12.85 7.71 15.14a94.28 94.28 0 0 0 9.9-11.12c2.28-3 5 1.85 4.91 4.34l-.42 20.72c-.08 3.87-1.57 4.55-2.63 4.55c-1.75 0-2.69-1.13-3-4c-.56-4.62-.1-14.9-.1-14.9L27.33 29c-3.39 3-5.19-6.32-7.79-9.49c-1.8 7.71-3.33 16-6.75 21.07c-.73 1.07-8.42 5.02-6.91 1.47Z"/>`),
		g.Group(children),
	)
}