package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BulletinNotice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd" clip-rule="evenodd"><path d="M18 6.923H2v11h16v-11Zm-16-1a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-11a1 1 0 0 0-1-1H2Z"/><path d="M6 9.423a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1-.5-.5Zm-1 3a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1h-9a.5.5 0 0 1-.5-.5Zm1 3a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1-.5-.5Zm5.768-13.025a2.5 2.5 0 0 0-3.536 0L4.354 6.277a.5.5 0 1 1-.708-.708l3.88-3.878a3.5 3.5 0 0 1 4.949 0l3.879 3.878a.5.5 0 1 1-.708.708l-3.878-3.88Z"/></g>`),
		g.Group(children),
	)
}