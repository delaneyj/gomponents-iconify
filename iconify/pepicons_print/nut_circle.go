package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NutCircle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><g opacity=".2"><path d="M10.58 7h6.84a1 1 0 0 1 .868.504l3.428 6a1 1 0 0 1 0 .992l-3.428 6a1 1 0 0 1-.868.504h-6.84a1 1 0 0 1-.868-.504l-3.428-6a1 1 0 0 1 0-.992l3.428-6A1 1 0 0 1 10.58 7Z"/><path fill-rule="evenodd" d="m11.16 9l-2.857 5l2.858 5h5.678l2.858-5l-2.858-5h-5.678Zm6.26-2h-6.84a1 1 0 0 0-.868.504l-3.428 6a1 1 0 0 0 0 .992l3.428 6a1 1 0 0 0 .868.504h6.84a1 1 0 0 0 .868-.504l3.428-6a1 1 0 0 0 0-.992l-3.428-6A1 1 0 0 0 17.42 7Z" clip-rule="evenodd"/></g><path fill-rule="evenodd" d="M17.722 19.744a1.5 1.5 0 0 1-1.302.756H9.58a1.5 1.5 0 0 1-1.302-.756l-3.429-6a1.5 1.5 0 0 1 0-1.488l3.429-6A1.5 1.5 0 0 1 9.58 5.5h6.84a1.5 1.5 0 0 1 1.302.756l3.429 6a1.5 1.5 0 0 1 0 1.488l-3.429 6ZM16.42 19.5a.5.5 0 0 0 .434-.252l3.428-6a.5.5 0 0 0 0-.496l-3.428-6a.5.5 0 0 0-.434-.252H9.58a.5.5 0 0 0-.434.252l-3.428 6a.5.5 0 0 0 0 .496l3.428 6a.5.5 0 0 0 .434.252h6.84Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 10.5a2.5 2.5 0 1 0 0 5a2.5 2.5 0 0 0 0-5ZM9.5 13a3.5 3.5 0 1 1 7 0a3.5 3.5 0 0 1-7 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}