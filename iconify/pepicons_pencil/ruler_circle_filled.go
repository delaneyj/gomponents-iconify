package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilRulerCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><path fill="#000" fill-rule="evenodd" d="M5.31 16.626a.5.5 0 0 0 0 .707l3.536 3.536a.5.5 0 0 0 .707 0L20.867 9.555a.5.5 0 0 0 0-.707l-3.536-3.535a.5.5 0 0 0-.707 0l-1.06 1.06l1.709 1.71a.5.5 0 1 1-.708.706L14.857 7.08l-1.415 1.415l.884.884a.5.5 0 0 1-.707.707l-.884-.884l-1.414 1.414l1.709 1.709a.5.5 0 1 1-.707.707l-1.709-1.709L9.2 12.737l.884.884a.5.5 0 1 1-.707.707l-.884-.884l-1.415 1.415l1.71 1.709a.5.5 0 1 1-.708.707l-1.709-1.71l-1.06 1.061Zm-.706 1.415a1.5 1.5 0 0 1 0-2.122L15.917 4.606a1.5 1.5 0 0 1 2.122 0l3.535 3.535a1.5 1.5 0 0 1 0 2.121L10.26 21.576a1.5 1.5 0 0 1-2.12 0l-3.536-3.535Z" clip-rule="evenodd"/></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilRulerCircleFilled0)"/></g>`),
		g.Group(children),
	)
}