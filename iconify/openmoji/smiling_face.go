package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmilingFace(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36" cy="36" r="23" fill="#FCEA2B"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><path d="M41.63 48.02c-1.123 1.267-3.049 2.078-5.78 2.078c-2.712 0-4.64-.802-5.776-2.054m-8.965-20.843a6.306 6.306 0 0 1 3.39-3a6.304 6.304 0 0 1 4.53-.42m21.69 3.42a7.19 7.19 0 0 0-7.91-3.43M23.484 34.245s3.932-2.17 8 0m9.25 0s3.932-2.17 8 0"/><circle cx="36" cy="36" r="23"/></g>`),
		g.Group(children),
	)
}