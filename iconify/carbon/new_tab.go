package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NewTab(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<defs><path id="carbonNewTab0" fill="currentColor" d="M26 26H6V6h10V4H6a2.002 2.002 0 0 0-2 2v20a2.002 2.002 0 0 0 2 2h20a2.002 2.002 0 0 0 2-2V16h-2Z"/></defs><use href="#carbonNewTab0"/><use href="#carbonNewTab0"/><path fill="currentColor" d="M26 6V2h-2v4h-4v2h4v4h2V8h4V6h-4z"/>`),
		g.Group(children),
	)
}