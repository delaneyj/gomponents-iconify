package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Eye(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<ellipse cx="35.75" cy="36.428" fill="#FFF" rx="34.81" ry="20.428"/><ellipse cx="35.75" cy="36.428" fill="#FFF" rx="34.81" ry="20.428"/><circle cx="36" cy="35.958" r="15.484" fill="#a57939"/><ellipse cx="35.75" cy="36.428" fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" rx="34.81" ry="20.428"/><circle cx="36" cy="35.958" r="8.442"/><circle cx="36" cy="35.958" r="8.442" fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"/><circle cx="36" cy="35.958" r="15.484" fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2"/>`),
		g.Group(children),
	)
}