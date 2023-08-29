package prime

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Image(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18 4.25H6A2.75 2.75 0 0 0 3.25 7v10A2.75 2.75 0 0 0 6 19.75h12A2.75 2.75 0 0 0 20.75 17V7A2.75 2.75 0 0 0 18 4.25ZM6 5.75h12A1.25 1.25 0 0 1 19.25 7v8.19l-2.72-2.72a.7.7 0 0 0-.56-.22a.79.79 0 0 0-.55.27l-1.29 1.55l-4.6-4.6A.7.7 0 0 0 9 9.25a.79.79 0 0 0-.55.27l-3.7 4.41V7A1.25 1.25 0 0 1 6 5.75ZM4.75 17v-.73l4.3-5.16l4.12 4.12l-2.52 3H6A1.25 1.25 0 0 1 4.75 17ZM18 18.25h-5.4l3.45-4.14l3.15 3.15a1.23 1.23 0 0 1-1.2.99Z"/>`),
		g.Group(children),
	)
}