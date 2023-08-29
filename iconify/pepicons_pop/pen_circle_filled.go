package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopPenCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M16.198 4.22L6.12 14.298a1 1 0 0 0-.282.555l-.705 4.594a1 1 0 0 0 1.14 1.14l4.595-.705a1 1 0 0 0 .555-.281L21.501 9.523a1 1 0 0 0 0-1.414l-3.89-3.89a1 1 0 0 0-1.413 0ZM7.317 18.404l.448-2.924l9.14-9.14l2.475 2.476l-9.14 9.14l-2.923.448Z" clip-rule="evenodd"/><path d="m14.442 8.247l1.06-1.061l3.242 3.24l-1.061 1.061l-3.241-3.24Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopPenCircleFilled0)"/></g>`),
		g.Group(children),
	)
}