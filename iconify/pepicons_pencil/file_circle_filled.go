package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilFileCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M9.5 15a.5.5 0 0 1 0-1h7a.5.5 0 0 1 0 1h-7Zm0 3a.5.5 0 0 1 0-1h7a.5.5 0 0 1 0 1h-7Z"/><path fill-rule="evenodd" d="M14.185 4H7.5A1.5 1.5 0 0 0 6 5.5v15A1.5 1.5 0 0 0 7.5 22h11a1.5 1.5 0 0 0 1.5-1.5V10.202a1.5 1.5 0 0 0-.395-1.014l-4.314-4.702A1.5 1.5 0 0 0 14.185 4ZM7 5.5a.5.5 0 0 1 .5-.5h6.685a.5.5 0 0 1 .369.162l4.314 4.702a.5.5 0 0 1 .132.338V20.5a.5.5 0 0 1-.5.5h-11a.5.5 0 0 1-.5-.5v-15Z" clip-rule="evenodd"/><path d="M14 10h5.5a.5.5 0 0 1 0 1h-6a.5.5 0 0 1-.5-.5v-6a.5.5 0 0 1 1 0V10Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilFileCircleFilled0)"/></g>`),
		g.Group(children),
	)
}