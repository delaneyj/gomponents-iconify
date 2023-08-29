package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScissorsCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13.5 26C20.404 26 26 20.404 26 13.5S20.404 1 13.5 1S1 6.596 1 13.5S6.596 26 13.5 26Zm0-2C19.299 24 24 19.299 24 13.5S19.299 3 13.5 3S3 7.701 3 13.5S7.701 24 13.5 24Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M9.5 22.5a3 3 0 0 0 2.435-4.753a.749.749 0 0 0 .595-.217l1.72-1.72l5.72 5.72a.75.75 0 1 0 1.06-1.06l-5.72-5.72l5.72-5.72a.75.75 0 0 0-1.06-1.06l-5.72 5.72l-1.72-1.72a.747.747 0 0 0-.316-.19a3 3 0 1 0-.869 1.085a.75.75 0 0 0 .125.165l1.72 1.72l-1.72 1.72a.748.748 0 0 0-.217.595A3 3 0 1 0 9.5 22.5Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M8.5 11.5a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-5a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm0 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-5a2 2 0 1 1 0 4a2 2 0 0 1 0-4Z" clip-rule="evenodd"/><path d="M19.978 18.782a.5.5 0 0 1-.697.718l-8.876-8.627a.5.5 0 1 1 .697-.717l8.876 8.626Z"/><path d="M10.146 16.146a.5.5 0 0 0 .708.708l9-9a.5.5 0 0 0-.708-.708l-9 9Z"/><path d="M4.15 4.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L4.151 4.878Z"/><path fill-rule="evenodd" d="M13 24.5c6.351 0 11.5-5.149 11.5-11.5S19.351 1.5 13 1.5S1.5 6.649 1.5 13S6.649 24.5 13 24.5Zm0 1c6.904 0 12.5-5.596 12.5-12.5S19.904.5 13 .5S.5 6.096.5 13S6.096 25.5 13 25.5Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}