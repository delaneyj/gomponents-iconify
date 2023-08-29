package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrCodeCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><g opacity=".2"><path d="M8 8.5a.5.5 0 0 1 .5-.5H21a.5.5 0 0 1 .5.5V21a.5.5 0 0 1-.5.5H8.5A.5.5 0 0 1 8 21V8.5Z"/><path fill-rule="evenodd" d="M9.5 9.5V20H20V9.5H9.5ZM8.5 8a.5.5 0 0 0-.5.5V21a.5.5 0 0 0 .5.5H21a.5.5 0 0 0 .5-.5V8.5A.5.5 0 0 0 21 8H8.5Z" clip-rule="evenodd"/></g><path fill-rule="evenodd" d="M8.5 8.5v3h3v-3h-3Zm-.5-1a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V8a.5.5 0 0 0-.5-.5H8Zm6.5 1v3h3v-3h-3Zm-.5-1a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5V8a.5.5 0 0 0-.5-.5h-4Zm-5.5 7v3h3v-3h-3Zm-.5-1a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5H8Z" clip-rule="evenodd"/><path d="M13.5 14a.5.5 0 0 1 .5-.5h4a.5.5 0 0 1 .5.5v4a.5.5 0 0 1-.5.5h-4a.5.5 0 0 1-.5-.5v-4Z"/><path fill-rule="evenodd" d="M14.5 14.5v3h3v-3h-3Zm-.5-1a.5.5 0 0 0-.5.5v4a.5.5 0 0 0 .5.5h4a.5.5 0 0 0 .5-.5v-4a.5.5 0 0 0-.5-.5h-4ZM5.5 6a.5.5 0 0 1 .5-.5h3.5a.5.5 0 0 1 0 1H6a.5.5 0 0 1-.5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6 5.5a.5.5 0 0 1 .5.5v3.5a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5Zm0 15a.5.5 0 0 1-.5-.5v-3.5a.5.5 0 0 1 1 0V20a.5.5 0 0 1-.5.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M5.5 20a.5.5 0 0 1 .5-.5h3.5a.5.5 0 0 1 0 1H6a.5.5 0 0 1-.5-.5Zm15 0a.5.5 0 0 1-.5.5h-3.5a.5.5 0 0 1 0-1H20a.5.5 0 0 1 .5.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M20 20.5a.5.5 0 0 1-.5-.5v-3.5a.5.5 0 0 1 1 0V20a.5.5 0 0 1-.5.5Zm0-15a.5.5 0 0 1 .5.5v3.5a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M20.5 6a.5.5 0 0 1-.5.5h-3.5a.5.5 0 0 1 0-1H20a.5.5 0 0 1 .5.5Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}