package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CpuWarning(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M19 2v4m0 4.01l.01-.011M16 14v1.4a.6.6 0 0 1-.6.6H8.6a.6.6 0 0 1-.6-.6V8.6a.6.6 0 0 1 .6-.6H14"/><path d="M20 14v5.4a.6.6 0 0 1-.6.6H4.6a.6.6 0 0 1-.6-.6V4.6a.6.6 0 0 1 .6-.6H14m6 13h2m-5 3v2m-5-2v2m-5-2v2m-3-5H2m2-5H2m2-5H2m10-3V2M7 4V2"/></g>`),
		g.Group(children),
	)
}