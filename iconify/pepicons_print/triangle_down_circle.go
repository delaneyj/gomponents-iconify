package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleDownCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" opacity=".2"/><path d="m14.75 17.014l3.778-6.514h-7.556l3.778 6.514Zm-.865 2.495a1 1 0 0 0 1.73 0l5.514-9.507a1 1 0 0 0-.865-1.502H9.236a1 1 0 0 0-.865 1.502l5.514 9.507Z" opacity=".2"/><path d="M13 19a.5.5 0 0 1-.429-.243l-6-10A.5.5 0 0 1 7 8h12a.5.5 0 0 1 .429.757l-6 10A.5.5 0 0 1 13 19Zm5.117-10H7.883L13 17.528L18.117 9Z"/><path d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z"/></g>`),
		g.Group(children),
	)
}