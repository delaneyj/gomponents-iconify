package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CommentO(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feCommentO0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feCommentO1" fill="currentColor" fill-rule="nonzero"><path id="feCommentO2" d="M5 21v-4.157c-1.25-1.46-2-3.319-2-5.343C3 6.806 7.03 3 12 3s9 3.806 9 8.5s-4.03 8.5-9 8.5a9.352 9.352 0 0 1-4.732-1.268L5 21Zm7-3c3.866 0 7-2.91 7-6.5S15.866 5 12 5s-7 2.91-7 6.5S8.134 18 12 18Z"/></g></g>`),
		g.Group(children),
	)
}