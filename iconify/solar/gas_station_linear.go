package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GasStationLinear(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="1.5"><path d="M16 22V8c0-2.828 0-4.243-.879-5.121C14.243 2 12.828 2 10 2H9c-2.828 0-4.243 0-5.121.879C3 3.757 3 5.172 3 8v14"/><path d="M11 6H8c-.943 0-1.414 0-1.707.293C6 6.586 6 7.057 6 8c0 .943 0 1.414.293 1.707C6.586 10 7.057 10 8 10h3c.943 0 1.414 0 1.707-.293C13 9.414 13 8.943 13 8c0-.943 0-1.414-.293-1.707C12.414 6 11.943 6 11 6Z"/><path stroke-linecap="round" d="M7 17h5m5 5H2M19.5 4l1.233.986c.138.11.207.166.27.222a3 3 0 0 1 .992 2.066c.005.084.005.172.005.348V18.5a1.5 1.5 0 0 1-3 0v-.071c0-.79-.64-1.429-1.429-1.429H16"/><path stroke-linecap="round" d="M22 8h-1.5A1.5 1.5 0 0 0 19 9.5v2.419a1.5 1.5 0 0 0 1.026 1.423L22 14"/></g>`),
		g.Group(children),
	)
}