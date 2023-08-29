package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func KissingFaceWithClosedEyes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36" cy="36" r="23" fill="#FCEA2B"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M36.415 41.092s10.525 3.196 0 5.564c0 0 10.458 2.988 0 4.913M21.108 27.202a6.306 6.306 0 0 1 3.391-3a6.304 6.304 0 0 1 4.53-.421m21.69 3.421a7.19 7.19 0 0 0-7.91-3.431M23.484 34.245s3.932-2.17 8 0m9.25 0s3.932-2.17 8 0"/><circle cx="36" cy="36" r="23"/></g>`),
		g.Group(children),
	)
}