package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ubq(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#00EA90"/><path fill="#FFF" fill-opacity=".698" fill-rule="nonzero" d="m18.215 7.508l7.777 4.068l-7.493 4.593l-.284-8.661zm-4.43 16.941l-7.777-4.068l7.493-4.594l.284 8.662z"/><path fill="#FFF" fill-rule="nonzero" d="M25.992 20.679L15.179 27v-8.869l10.813-6.555v9.103zm-19.984-9.4L16.821 5v8.834L6.008 20.381v-9.103z"/></g>`),
		g.Group(children),
	)
}