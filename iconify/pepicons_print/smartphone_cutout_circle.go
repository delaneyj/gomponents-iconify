package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartphoneCutoutCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" opacity=".2"/><path d="M9 5.5A1.5 1.5 0 0 1 10.5 4h8A1.5 1.5 0 0 1 20 5.5v16a1.5 1.5 0 0 1-1.5 1.5h-8A1.5 1.5 0 0 1 9 21.5v-16Z" opacity=".2"/><path d="M7 5a1.5 1.5 0 0 1 1.5-1.5h9A1.5 1.5 0 0 1 19 5v16a1.5 1.5 0 0 1-1.5 1.5h-9A1.5 1.5 0 0 1 7 21V5Zm1.5-.5A.5.5 0 0 0 8 5v16a.5.5 0 0 0 .5.5h9a.5.5 0 0 0 .5-.5V5a.5.5 0 0 0-.5-.5h-9Z"/><path d="M10.5 6a1 1 0 0 1 1-1h3a1 1 0 1 1 0 2h-3a1 1 0 0 1-1-1Z"/><path d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z"/></g>`),
		g.Group(children),
	)
}