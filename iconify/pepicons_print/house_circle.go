package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HouseCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><path d="M6.889 14H5.5a.5.5 0 0 1-.33-.875l8.5-7.5a.5.5 0 0 1 .66 0l8.5 7.5a.5.5 0 0 1-.33.875h-1.389v7a.5.5 0 0 1-.5.5H7.39a.5.5 0 0 1-.5-.5v-7Z" opacity=".2"/><path fill-rule="evenodd" d="M4 13h1.389v7a.5.5 0 0 0 .5.5H19.11a.5.5 0 0 0 .5-.5v-7H21a.5.5 0 0 0 .33-.875l-8.5-7.5a.5.5 0 0 0-.66 0l-8.5 7.5A.5.5 0 0 0 4 13Zm1.889-1h-.567L12.5 5.667L19.678 12h-.567a.5.5 0 0 0-.5.5v7H6.39v-7a.5.5 0 0 0-.5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13.708 14.5h-2.5a1 1 0 0 0-1 1v4a1 1 0 0 0 1 1h2.5a1 1 0 0 0 1-1v-4a1 1 0 0 0-1-1Zm-2.5 5v-4h2.5v4h-2.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}