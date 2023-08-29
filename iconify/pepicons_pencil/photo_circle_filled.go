package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PhotoCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilPhotoCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M11 11a1 1 0 1 0 0-2a1 1 0 0 0 0 2Zm0 1a2 2 0 1 0 0-4a2 2 0 0 0 0 4Z" clip-rule="evenodd"/><path d="M10.5 13.678L8.59 16.25h1.91l1.91 1H8.59a1 1 0 0 1-.802-1.596l1.909-2.572a1 1 0 0 1 1.606 0l.955 1.286l-.803.596l-.955-1.286Z"/><path fill-rule="evenodd" d="m17.257 16.25l-3.007-3.472l-3.007 3.472h6.014Zm-2.251-4.127a1 1 0 0 0-1.512 0l-3.007 3.472a1 1 0 0 0 .756 1.655h6.014a1 1 0 0 0 .756-1.655l-3.007-3.472Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M20 5.5H6a.5.5 0 0 0-.5.5v14a.5.5 0 0 0 .5.5h14a.5.5 0 0 0 .5-.5V6a.5.5 0 0 0-.5-.5Zm-14-1A1.5 1.5 0 0 0 4.5 6v14A1.5 1.5 0 0 0 6 21.5h14a1.5 1.5 0 0 0 1.5-1.5V6A1.5 1.5 0 0 0 20 4.5H6Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilPhotoCircleFilled0)"/></g>`),
		g.Group(children),
	)
}