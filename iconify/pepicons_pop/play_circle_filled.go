package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PlayCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopPlayCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" d="M18.544 12.59a1 1 0 0 1-.053 1.728L9.476 19.2A1 1 0 0 1 8 18.321V7.804a1 1 0 0 1 1.53-.848l9.014 5.634Z"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopPlayCircleFilled0)"/></g>`),
		g.Group(children),
	)
}