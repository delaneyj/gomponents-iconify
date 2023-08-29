package ri

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScalesFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12.998 2v1h7v2h-7v14h4v2h-10v-2h4V5h-7V3h7V2h2Zm-8 4.343l2.828 2.829a4 4 0 1 1-5.657 0L5 6.343Zm14 0l2.828 2.829a4 4 0 1 1-5.657 0l2.829-2.829Zm0 2.829l-1.414 1.414A1.987 1.987 0 0 0 16.998 12l4 .001c0-.54-.212-1.041-.586-1.415l-1.414-1.414Zm-14 0l-1.414 1.414A1.987 1.987 0 0 0 2.998 12l4 .001c0-.54-.212-1.041-.586-1.415L4.998 9.172Z"/>`),
		g.Group(children),
	)
}