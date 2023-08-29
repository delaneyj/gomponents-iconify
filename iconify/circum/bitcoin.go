package circum

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bitcoin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M12 21.934A9.934 9.934 0 1 1 21.934 12A9.945 9.945 0 0 1 12 21.934Zm0-18.868A8.934 8.934 0 1 0 20.934 12A8.944 8.944 0 0 0 12 3.066Z"/><path fill="currentColor" d="M14.28 11.78a1.994 1.994 0 0 0-1.53-3.28h-.25V7.47A.489.489 0 0 0 12 7a.483.483 0 0 0-.5.47V8.5h-1.25a1 1 0 0 0-1 1v5a1 1 0 0 0 1 1h1.25v1.03a.483.483 0 0 0 .5.47a.489.489 0 0 0 .5-.47V15.5h.75a2.006 2.006 0 0 0 2-2a2.033 2.033 0 0 0-.97-1.72ZM10.25 9.5h2.5a1 1 0 0 1 0 2h-2.5Zm3 5h-3v-2h3a1 1 0 0 1 0 2Z"/>`),
		g.Group(children),
	)
}