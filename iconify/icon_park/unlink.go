package icon_park

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Unlink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><path d="M25.8927 16.0307L18.1145 8.2525C15.2506 5.38866 10.7031 5.29302 7.9572 8.0389C5.21132 10.7848 5.30696 15.3323 8.1708 18.1962L15.949 25.9744"/><path d="M31.9161 22.0707L39.6943 29.8489C42.5581 32.7127 42.9291 37.1233 39.9079 40.0062C36.8867 42.8891 32.6144 42.6564 29.7506 39.7926L21.9724 32.0144"/><path d="M21.2384 21.0759L17.3493 17.1868"/><path d="M30.3131 30.1504L26.424 26.2613"/></g>`),
		g.Group(children),
	)
}