package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FolderCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><path d="M15.219 10.5H20.5A2.5 2.5 0 0 1 23 13v7a2.5 2.5 0 0 1-2.5 2.5h-11A2.5 2.5 0 0 1 7 20v-9.5A2.5 2.5 0 0 1 9.5 8h2.84a2.5 2.5 0 0 1 2.17 1.26l.709 1.24Z" opacity=".2"/><path fill-rule="evenodd" d="M18.5 8.5h-5.281l-.709-1.24A2.5 2.5 0 0 0 10.34 6H7.5A2.5 2.5 0 0 0 5 8.5V18a2.5 2.5 0 0 0 2.5 2.5h11A2.5 2.5 0 0 0 21 18v-7a2.5 2.5 0 0 0-2.5-2.5Zm-11 11A1.5 1.5 0 0 1 6 18V8.5A1.5 1.5 0 0 1 7.5 7h2.84a1.5 1.5 0 0 1 1.302.756l.852 1.492a.5.5 0 0 0 .435.252H18.5A1.5 1.5 0 0 1 20 11v7a1.5 1.5 0 0 1-1.5 1.5h-11Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}