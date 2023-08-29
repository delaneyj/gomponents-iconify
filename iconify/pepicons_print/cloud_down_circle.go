package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CloudDownCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M16 7h-1a4.002 4.002 0 0 0-3.874 3H11a4 4 0 1 0 0 8h8a4 4 0 0 0 .899-7.899A4.002 4.002 0 0 0 16 7Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M14 6h-1a4.002 4.002 0 0 0-3.874 3H9a4 4 0 1 0 0 8h8a4 4 0 0 0 .899-7.899A4.002 4.002 0 0 0 14 6Zm-4.099 4l.193-.75A3.002 3.002 0 0 1 13 7h1c1.405 0 2.614.975 2.924 2.325l.14.61l.61.141A3.001 3.001 0 0 1 17 16H9a3 3 0 1 1 0-6h.901Z" clip-rule="evenodd"/><path d="M13 13a.5.5 0 0 1 1 0v7.5a.5.5 0 0 1-1 0V13Z"/><path d="M15.688 18.11a.5.5 0 0 1 .624.78l-2.5 2a.5.5 0 1 1-.624-.78l2.5-2Z"/><path d="M10.688 18.89a.5.5 0 0 1 .624-.78l2.5 2a.5.5 0 1 1-.624.78l-2.5-2Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}