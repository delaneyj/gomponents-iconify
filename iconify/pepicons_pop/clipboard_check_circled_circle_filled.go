package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardCheckCircledCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopClipboardCheckCircledCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M8.68 5.5a1 1 0 0 1 1-1h6.643a1 1 0 0 1 1 1v3.875a1 1 0 0 1-1 1H9.68a1 1 0 0 1-1-1V5.5Zm2 1v1.875h4.643V6.5H10.68Z"/><path d="M8 19V8h1V6H7a1 1 0 0 0-1 1v13a1 1 0 0 0 1 1h6.337a5.526 5.526 0 0 1-1.737-2H8Zm10-7.793c.742.21 1.421.572 2 1.05V7a1 1 0 0 0-1-1h-2v2h1v3.207Z"/><path d="M16.5 20a3.5 3.5 0 1 0 0-7a3.5 3.5 0 0 0 0 7Zm0 2a5.5 5.5 0 1 0 0-11a5.5 5.5 0 0 0 0 11Z"/><path d="M18.87 14.72a1 1 0 0 1 .156 1.405l-1.65 2.063a1.5 1.5 0 0 1-2.233.124l-1.104-1.105a1 1 0 0 1 1.414-1.414l.71.71l1.302-1.628a1 1 0 0 1 1.405-.156Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopClipboardCheckCircledCircleFilled0)"/></g>`),
		g.Group(children),
	)
}