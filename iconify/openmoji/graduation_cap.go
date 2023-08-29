package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GraduationCap(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#3F3F3F" stroke="#3F3F3F" stroke-miterlimit="10" stroke-width="2" d="M56 27.917v17.979H16V27.917"/><path fill="#F1B31C" stroke="#F1B31C" stroke-miterlimit="10" stroke-width="2" d="M8.793 28.5v8.406"/><circle cx="8.793" cy="36.906" r="2" fill="#F1B31C"/><path fill="#F1B31C" stroke="#F1B31C" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="1.747" d="M10.793 45.896H7l1.897-8.12h0z"/><path fill="#3F3F3F" d="M4 22.875h64v5.042H4zm52 8.042v14.979H16V30.917"/><circle cx="36" cy="22.002" r="3"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M4 22.875h64v5.042H4z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="1.747" d="M56 30.917v14.979H16V30.917"/>`),
		g.Group(children),
	)
}