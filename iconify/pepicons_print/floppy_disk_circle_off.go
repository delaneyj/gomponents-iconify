package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FloppyDiskCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><g opacity=".2"><path d="M7.75 21.5h13a1 1 0 0 0 1-1V9.914a1 1 0 0 0-.293-.707l-2.414-2.414a1 1 0 0 0-.707-.293H7.75a1 1 0 0 0-1 1v13a1 1 0 0 0 1 1Z"/><path fill-rule="evenodd" d="M19.75 19.5v-9.172L17.922 8.5H8.75v11h11Zm1 2h-13a1 1 0 0 1-1-1v-13a1 1 0 0 1 1-1h10.586a1 1 0 0 1 .707.293l2.414 2.414a1 1 0 0 1 .293.707V20.5a1 1 0 0 1-1 1Z" clip-rule="evenodd"/></g><path fill-rule="evenodd" d="M6.5 5.75a.5.5 0 0 0-.5.5v13a.5.5 0 0 0 .5.5h13a.5.5 0 0 0 .5-.5V8.664a.5.5 0 0 0-.146-.353l-2.415-2.415a.5.5 0 0 0-.353-.146H6.5Zm-1.5.5a1.5 1.5 0 0 1 1.5-1.5h10.586a1.5 1.5 0 0 1 1.06.44l2.415 2.414A1.5 1.5 0 0 1 21 8.664V19.25a1.5 1.5 0 0 1-1.5 1.5h-13a1.5 1.5 0 0 1-1.5-1.5v-13Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M8.5 13.75a1.5 1.5 0 0 1 1.5-1.5h6a1.5 1.5 0 0 1 1.5 1.5v7h-1v-7a.5.5 0 0 0-.5-.5h-6a.5.5 0 0 0-.5.5v7h-1v-7Zm.5-6a1.5 1.5 0 0 0 1.5 1.5h4a1.5 1.5 0 0 0 1.5-1.5v-2a.5.5 0 0 0-1 0v2a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5v-2a.5.5 0 0 0-1 0v2Z" clip-rule="evenodd"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}