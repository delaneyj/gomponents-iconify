package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MusicalNotes(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g fill="#3f3f3f"><circle cx="15.205" cy="45.71" r="5.548"/><circle cx="29.446" cy="57.93" r="5.548"/><circle cx="46.139" cy="45.71" r="5.548"/></g><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"><circle cx="15.205" cy="45.71" r="5"/><circle cx="29.446" cy="57.93" r="5"/><circle cx="46.139" cy="45.71" r="5"/><path d="M56.729 17.155h-5.968l.378 28.555M40.036 29.375h-5.969l.379 28.555m-8.651-40.775h-5.969l.379 28.555"/></g>`),
		g.Group(children),
	)
}