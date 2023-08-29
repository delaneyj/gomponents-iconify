package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StudioBackdropCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopStudioBackdropCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M6 6a1.5 1.5 0 0 1 1.5-1.5h11A1.5 1.5 0 0 1 20 6v7a1.5 1.5 0 0 1-1.5 1.5h-11A1.5 1.5 0 0 1 6 13V6Zm2 .5v6h10v-6H8Z"/><path d="M4.75 5.5a1 1 0 0 1 1-1h14.5a1 1 0 1 1 0 2H5.75a1 1 0 0 1-1-1Zm1.97 14l1.26-6.304l-1.96-.392l-1.381 6.902a1.5 1.5 0 0 0 1.47 1.794H19.89a1.5 1.5 0 0 0 1.471-1.794l-1.38-6.902l-1.962.392L19.28 19.5H6.72Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopStudioBackdropCircleFilled0)"/></g>`),
		g.Group(children),
	)
}