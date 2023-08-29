package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShareAndroidCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M8 15.5a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0-4a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm9-1a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0-4a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Zm0 14a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5Zm0-4a1.5 1.5 0 1 1 0 3a1.5 1.5 0 0 1 0-3Z" clip-rule="evenodd"/><path d="m9.754 12.18l-.508-.86l5.5-3.25l.508.86l-5.5 3.25ZM15 17.878l.479-.878l-5.5-3l-.479.878l5.5 3Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}