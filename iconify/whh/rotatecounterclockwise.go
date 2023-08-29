package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Rotatecounterclockwise(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960 704H832q-26 0-45-19t-19-45V64q0-26 19-45t45-19h128q27 0 45.5 19t18.5 45v576q0 27-18.5 45.5T960 704zm-320 320H64q-26 0-45-19T0 960V832q0-27 19-45.5T64 768h576q27 0 45.5 18.5T704 832v128q0 26-18.5 45t-45.5 19zm-64-192H128q-26 0-45 18.5T64 896t19 45.5t45 18.5h448q27 0 45.5-18.5T640 896t-18.5-45.5T576 832zm0-576h-64q-89 0-157.5 54.5T264 448h109q11 10 11 23t-11 23L218 694q-11 10-26 10t-26-10L11 494Q0 484 0 471t11-23h123q23-137 129.5-228.5T512 128h64q27 0 45.5 18.5T640 192t-18.5 45.5T576 256z"/>`),
		g.Group(children),
	)
}