package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PersonCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><path fill="currentColor" fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><g transform="translate(3 3)"><g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd" opacity=".2"><path d="M7.75 6.6c0-1.95 1.53-3.6 3.5-3.6s3.5 1.65 3.5 3.6c0 1.95-1.53 3.6-3.5 3.6s-3.5-1.65-3.5-3.6Z"/><path d="M11.264 9.067c-3.225 0-6.014 2.471-6.014 5.766l.002 2.168A1 1 0 0 0 6.25 18h10a1 1 0 0 0 1-1v-2.167c0-3.288-2.755-5.766-5.986-5.766Z"/></g><circle cx="9.5" cy="5.5" r="3" stroke="currentColor" stroke-linecap="round"/><path stroke="currentColor" stroke-linecap="round" d="M15 16.5v-2c0-3.098-2.495-6-5.5-6c-3.006 0-5.5 2.902-5.5 6v2"/></g><path fill="currentColor" fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}