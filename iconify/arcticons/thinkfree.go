package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Thinkfree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.28 5.5a5.61 5.61 0 0 0-1.8.29l-.84.3V42l.78.29a4.23 4.23 0 0 0 1.92.25c1.13 0 1.2-.06 6-2l7.94-3.26a27.37 27.37 0 0 0 3.41-1.57a3.17 3.17 0 0 0 .58-.8c.25-.47.26-.94.26-10.81c0-11.23 0-10.8-.77-11.55c-.38-.36-3-1.49-13.22-5.63c-3.1-1.34-3.36-1.42-4.26-1.42Zm-5.4 1.97v33.15m-5.36-30.46L13.39 38M8 35.36L7.3 35a3.81 3.81 0 0 1-1-.8c-.79-1-.78-.83-.78-10.07c0-8.11 0-8.5.27-9.16a3.29 3.29 0 0 1 1.51-1.7l.64-.27L8 35.36Zm-.06-22.4l13.7-6.87M7.97 35.36l13.67 6.59m5.26-7.39l13.6-3.25m-5.72 1.37V14.1M27 11.96v16.45m7.78-14.31L26.9 28.41m13.6-12.74v11.17m0-11.17L27 11.96m7.78 2.14l5.72 12.74"/>`),
		g.Group(children),
	)
}