package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SignalStrength(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTSignalStrength0"><g fill="none"><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M38.142 38.142c7.81-7.81 7.81-20.474 0-28.284c-7.81-7.81-20.474-7.81-28.284 0c-7.81 7.81-7.81 20.474 0 28.284m22.627-5.657c4.687-4.686 4.687-12.284 0-16.97c-4.686-4.687-12.284-4.687-16.97 0c-4.687 4.686-4.687 12.284 0 16.97"/><path fill="#555" d="M28 24a4 4 0 1 1-8 0a4 4 0 0 1 8 0Z"/><path stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4" d="M24 28a4 4 0 1 0 0-8a4 4 0 0 0 0 8Zm0 0v16m0 0h4m-4 0h-4"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTSignalStrength0)"/>`),
		g.Group(children),
	)
}