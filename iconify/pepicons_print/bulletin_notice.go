package pepicons_print

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BulletinNotice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><g opacity=".2"><path d="M2 8a1 1 0 0 1 1-1h16a1 1 0 0 1 1 1v11a1 1 0 0 1-1 1H3a1 1 0 0 1-1-1V8Z"/><path fill-rule="evenodd" d="M4 9v9h14V9H4ZM3 7a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1V8a1 1 0 0 0-1-1H3Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M12.414 3.828a2 2 0 0 0-2.828 0l-3.879 3.88a1 1 0 1 1-1.414-1.415l3.879-3.879a4 4 0 0 1 5.656 0l3.88 3.879a1 1 0 0 1-1.415 1.414l-3.879-3.879Z" clip-rule="evenodd"/></g><path fill-rule="evenodd" d="M18 6.923H2v11h16v-11Zm-16-1a1 1 0 0 0-1 1v11a1 1 0 0 0 1 1h16a1 1 0 0 0 1-1v-11a1 1 0 0 0-1-1H2Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M6 9.423a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1-.5-.5Zm-1 3a.5.5 0 0 1 .5-.5h9a.5.5 0 0 1 0 1h-9a.5.5 0 0 1-.5-.5Zm1 3a.5.5 0 0 1 .5-.5h7a.5.5 0 0 1 0 1h-7a.5.5 0 0 1-.5-.5Zm5.768-13.025a2.5 2.5 0 0 0-3.536 0L4.354 6.277a.5.5 0 1 1-.708-.708l3.88-3.878a3.5 3.5 0 0 1 4.949 0l3.879 3.878a.5.5 0 1 1-.708.708l-3.878-3.88Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}