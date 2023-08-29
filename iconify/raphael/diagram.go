package raphael

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Diagram(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m6.812 17.202l7.396-3.665v-2.164h-.834c-.414 0-.808-.084-1.167-.237v1.16L4.81 15.96v2.912h2V17.2zm19.75 1.673v-2.913l-7.397-3.666v-1.158a2.98 2.98 0 0 1-1.166.236h-.833v2.164l7.395 3.666v1.672h2zm-9.874 0v-7.5h-2v7.5h2zm11.187 1H23.25a2 2 0 0 0-2 2V26.5a2 2 0 0 0 2 2h4.625a2 2 0 0 0 2-2v-4.625a2 2 0 0 0-2-2zm-19.75 0H3.5a2 2 0 0 0-2 2V26.5a2 2 0 0 0 2 2h4.625a2 2 0 0 0 2-2v-4.625a2 2 0 0 0-2-2zm5.25-9.5H18a2 2 0 0 0 2-2V3.75a2 2 0 0 0-2-2h-4.625a2 2 0 0 0-2 2v4.625a2 2 0 0 0 2 2zm4.625 9.5h-4.625a2 2 0 0 0-2 2V26.5a2 2 0 0 0 2 2H18a2 2 0 0 0 2-2v-4.625a2 2 0 0 0-2-2z"/>`),
		g.Group(children),
	)
}