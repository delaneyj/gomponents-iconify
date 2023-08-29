package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EndpointDisplacement(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTEndpointDisplacement0"><g fill="none"><path fill="#555" d="M16 10a5 5 0 1 1-10 0a5 5 0 0 1 10 0Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M16 10a5 5 0 1 1-10 0a5 5 0 0 1 10 0Zm0 0h11m0 0l-4-4m4 4l-4 4"/><path fill="#555" d="M32 38a5 5 0 1 0 10 0a5 5 0 0 0-10 0Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M32 38a5 5 0 1 0 10 0a5 5 0 0 0-10 0Zm0 0H21m0 0l4-4m-4 4l4 4"/><path fill="#555" d="M33 11a5 5 0 1 0 10 0a5 5 0 0 0-10 0Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38 16a5 5 0 1 1 0-10a5 5 0 0 1 0 10Zm0 0v11m0 0l4-4m-4 4l-4-4"/><path fill="#555" d="M5 37a5 5 0 1 1 10 0a5 5 0 0 1-10 0Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M10 32a5 5 0 1 0 0 10a5 5 0 0 0 0-10Zm0 0V21m0 0l4 4m-4-4l-4 4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTEndpointDisplacement0)"/>`),
		g.Group(children),
	)
}