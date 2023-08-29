package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LockOpenCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M12.5 12.5H17a3.5 3.5 0 0 1 3.5 3.5v3a3.5 3.5 0 0 1-3.5 3.5h-5A3.5 3.5 0 0 1 8.5 19v-3a3.502 3.502 0 0 1 2.5-3.355V9.75C11 7.693 12.552 6 14.5 6S18 7.693 18 9.75V10a.75.75 0 0 1-1.5 0v-.25c0-1.257-.91-2.25-2-2.25c-1.09 0-2 .993-2 2.25v2.75Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M10.5 15.5a2 2 0 1 0 4 0a2 2 0 0 0-4 0Zm3 0a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M15 10.5h-5A3.5 3.5 0 0 0 6.5 14v3a3.5 3.5 0 0 0 3.5 3.5h5a3.5 3.5 0 0 0 3.5-3.5v-3a3.5 3.5 0 0 0-3.5-3.5ZM7.5 14a2.5 2.5 0 0 1 2.5-2.5h5a2.5 2.5 0 0 1 2.5 2.5v3a2.5 2.5 0 0 1-2.5 2.5h-5A2.5 2.5 0 0 1 7.5 17v-3Z" clip-rule="evenodd"/><path d="M10 11a.5.5 0 0 1-1 0V7.5a3.5 3.5 0 1 1 7 0v1a.5.5 0 0 1-1 0v-1a2.5 2.5 0 0 0-5 0V11Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}