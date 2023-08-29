package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StudioBackdropCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><g opacity=".2"><path d="M8.25 6.5a.5.5 0 0 1 .5-.5h11a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-.5.5h-11a.5.5 0 0 1-.5-.5v-7Z"/><path fill-rule="evenodd" d="M10.25 8v4h8V8h-8Zm-1.5-2a.5.5 0 0 0-.5.5v7a.5.5 0 0 0 .5.5h11a.5.5 0 0 0 .5-.5v-7a.5.5 0 0 0-.5-.5h-11Z" clip-rule="evenodd"/><path d="M8.25 13.5a.5.5 0 0 1 .5-.5h11a.5.5 0 0 1 .5.5v7a.5.5 0 0 1-.5.5h-11a.5.5 0 0 1-.5-.5v-7Z"/><path fill-rule="evenodd" d="M10.25 15v4h8v-4h-8Zm-1.5-2a.5.5 0 0 0-.5.5v7a.5.5 0 0 0 .5.5h11a.5.5 0 0 0 .5-.5v-7a.5.5 0 0 0-.5-.5h-11Zm-2.5-7A.75.75 0 0 1 7 5.25h14.5a.75.75 0 0 1 0 1.5H7A.75.75 0 0 1 6.25 6Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="m8.576 19.5l1.636-5.725l-1.924-.55l-1.908 6.682a1.25 1.25 0 0 0 1.2 1.593h13.34c.83 0 1.43-.795 1.201-1.593l-1.909-6.682l-1.922.55l1.635 5.725H8.576Z" clip-rule="evenodd"/></g><path fill-rule="evenodd" d="M6.5 6a1 1 0 0 1 1-1h11a1 1 0 0 1 1 1v7a1 1 0 0 1-1 1h-11a1 1 0 0 1-1-1V6Zm12 0h-11v7h11V6Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M5.25 5.5a.5.5 0 0 1 .5-.5h14.5a.5.5 0 0 1 0 1H5.75a.5.5 0 0 1-.5-.5Zm2.24 7.598L6.11 20h13.78l-1.38-6.902l.98-.196l1.38 6.902A1 1 0 0 1 19.89 21H6.11a1 1 0 0 1-.98-1.196l1.38-6.902l.98.196Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}