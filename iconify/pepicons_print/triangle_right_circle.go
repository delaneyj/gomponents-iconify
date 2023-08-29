package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TriangleRightCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" opacity=".2"/><path d="M17.514 14.25L11 10.472v7.556l6.514-3.778Zm2.495.865a1 1 0 0 0 0-1.73l-9.507-5.514A1 1 0 0 0 9 8.736v11.028a1 1 0 0 0 1.502.865l9.507-5.514Z" opacity=".2"/><path d="M19 13a.5.5 0 0 1-.243.429l-10 6A.5.5 0 0 1 8 19V7a.5.5 0 0 1 .757-.429l10 6A.5.5 0 0 1 19 13ZM9 7.883v10.234L17.528 13L9 7.883Z"/><path d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z"/></g>`),
		g.Group(children),
	)
}