package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircleBigCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" opacity=".2"/><path d="M13.5 23a9.5 9.5 0 1 0 0-19a9.5 9.5 0 0 0 0 19Zm0-2.111a7.389 7.389 0 1 0 0-14.778a7.389 7.389 0 0 0 0 14.778Z" opacity=".2"/><path d="M13 4.5a8.5 8.5 0 1 0 0 17a8.5 8.5 0 0 0 0-17ZM3.5 13a9.5 9.5 0 1 1 19 0a9.5 9.5 0 0 1-19 0Z"/><path d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z"/></g>`),
		g.Group(children),
	)
}