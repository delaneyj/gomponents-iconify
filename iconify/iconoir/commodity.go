package iconoir

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Commodity(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-width="1.5"><path d="m12.147 18.28l1.184-5.8a.6.6 0 0 1 .588-.48h6.162a.6.6 0 0 1 .588.48l1.184 5.8a.6.6 0 0 1-.588.72h-8.53a.6.6 0 0 1-.588-.72Z"/><path d="m7.147 11.28l1.184-5.8A.6.6 0 0 1 8.918 5h6.164a.6.6 0 0 1 .587.48l1.184 5.8a.6.6 0 0 1-.588.72h-8.53a.6.6 0 0 1-.588-.72Z"/><path d="m2.147 18.28l1.184-5.8a.6.6 0 0 1 .587-.48h6.163a.6.6 0 0 1 .588.48l1.184 5.8a.6.6 0 0 1-.588.72h-8.53a.6.6 0 0 1-.588-.72Z"/></g>`),
		g.Group(children),
	)
}