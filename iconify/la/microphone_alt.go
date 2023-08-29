package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrophoneAlt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M21 4c-3.578 0-6.531 2.715-6.938 6.188L6 21.593l-.469.687l.563.594l.812.813l-2.625 2.593l1.439 1.44l2.562-2.594L9.656 26.5l.719-.5l11.469-8.063C25.304 17.52 28 14.57 28 11c0-3.855-3.145-7-7-7zm0 2c2.773 0 5 2.227 5 5a5 5 0 0 1-.813 2.75L18.25 6.812A5 5 0 0 1 21 6zm-4.188 2.25l6.938 6.938A5 5 0 0 1 21 16c-2.773 0-5-2.227-5-5a5 5 0 0 1 .813-2.75zm-2.437 4.938a7.068 7.068 0 0 0 4.406 4.437l-8.875 6.281l-1.781-1.843z"/>`),
		g.Group(children),
	)
}