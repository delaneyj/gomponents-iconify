package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScissorsOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M6.5 19.5a3 3 0 0 0 2.435-4.753a.749.749 0 0 0 .595-.217l1.72-1.72l5.72 5.72a.75.75 0 1 0 1.06-1.06l-5.72-5.72l5.72-5.72a.75.75 0 0 0-1.06-1.06l-5.72 5.72l-1.72-1.72a.747.747 0 0 0-.316-.19a3 3 0 1 0-.869 1.085a.75.75 0 0 0 .125.165l1.72 1.72l-1.72 1.72a.748.748 0 0 0-.217.595A3 3 0 1 0 6.5 19.5Z" clip-rule="evenodd" opacity=".2"/><path fill-rule="evenodd" d="M5.5 8.5a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-5a2 2 0 1 1 0 4a2 2 0 0 1 0-4Zm0 15a3 3 0 1 0 0-6a3 3 0 0 0 0 6Zm0-5a2 2 0 1 1 0 4a2 2 0 0 1 0-4Z" clip-rule="evenodd"/><path d="M16.978 15.782a.5.5 0 0 1-.697.718L7.405 7.873a.5.5 0 1 1 .697-.717l8.876 8.626Z"/><path d="M7.146 13.146a.5.5 0 0 0 .708.708l9-9a.5.5 0 0 0-.708-.708l-9 9Z"/><path d="M1.15 1.878a.514.514 0 0 1 .728-.727l16.971 16.971a.514.514 0 0 1-.727.727L1.151 1.878Z"/></g>`),
		g.Group(children),
	)
}