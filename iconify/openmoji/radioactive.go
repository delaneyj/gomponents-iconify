package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Radioactive(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36" cy="36" r="23" fill="#f4aa41"/><g stroke="#000"><circle cx="35.96" cy="35.89" r="3.281" stroke-miterlimit="10" stroke-width="1.25"/><path stroke-miterlimit="10" stroke-width="1.254" d="M23.09 28.58a14.76 14.76 0 0 1 3.496-4.097l5.791 6.957a5.733 5.733 0 0 0-2.167 3.753l-8.921-1.537a14.76 14.76 0 0 1 1.8-5.076zm25.82 0a14.76 14.76 0 0 1 1.8 5.076l-8.921 1.537a5.733 5.733 0 0 0-2.167-3.753l5.791-6.957a14.76 14.76 0 0 1 3.496 4.097zM36 50.87a14.76 14.76 0 0 1-5.296-.98l3.129-8.493a5.733 5.733 0 0 0 4.333 0l3.129 8.494a14.76 14.76 0 0 1-5.296.979z"/><circle cx="36" cy="36" r="23" fill="none" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"/></g>`),
		g.Group(children),
	)
}