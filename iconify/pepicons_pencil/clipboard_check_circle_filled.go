package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardCheckCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilClipboardCheckCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M9.175 5.5a.5.5 0 0 1 .5-.5h6.643a.5.5 0 0 1 .5.5v3.875a.5.5 0 0 1-.5.5H9.675a.5.5 0 0 1-.5-.5V5.5Zm1 .5v2.875h5.643V6h-5.643Z" clip-rule="evenodd"/><path d="M7.5 8v12h11V8h-2V7h2a1 1 0 0 1 1 1v12a1 1 0 0 1-1 1h-11a1 1 0 0 1-1-1V8a1 1 0 0 1 1-1h2v1h-2Z"/><path fill-rule="evenodd" d="M15.284 12.088a.5.5 0 0 1 .128.696l-1.767 2.564a1 1 0 0 1-1.437.222l-1.515-1.175a.5.5 0 1 1 .614-.79l1.514 1.175l1.767-2.564a.5.5 0 0 1 .696-.128Z" clip-rule="evenodd"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilClipboardCheckCircleFilled0)"/></g>`),
		g.Group(children),
	)
}