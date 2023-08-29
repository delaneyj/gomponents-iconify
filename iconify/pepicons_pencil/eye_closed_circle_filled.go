package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func EyeClosedCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilEyeClosedCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path d="M5.094 11.014a.5.5 0 1 1 .812-.583c.348.484.87.934 1.536 1.318c1.373.792 3.25 1.251 5.262 1.251c2.012 0 3.89-.459 5.263-1.251c.665-.384 1.187-.834 1.535-1.318a.5.5 0 0 1 .813.583c-.437.608-1.067 1.15-1.848 1.601c-1.532.884-3.583 1.385-5.763 1.385c-2.18 0-4.23-.5-5.762-1.385c-.782-.451-1.412-.993-1.848-1.6Z"/><path d="M13.5 14a.5.5 0 0 0-1 0v2.5a.5.5 0 0 0 1 0V14Zm-4.49-.598a.5.5 0 1 1 .98.196l-.5 2.5a.5.5 0 0 1-.98-.196l.5-2.5Zm7.98 0a.5.5 0 0 0-.98.196l.5 2.5a.5.5 0 0 0 .98-.196l-.5-2.5Zm2.364-1.756a.5.5 0 0 0-.708.708l2 2a.5.5 0 0 0 .708-.708l-2-2Zm-12.906.018a.5.5 0 1 1 .74.672l-1.818 2a.5.5 0 0 1-.74-.672l1.818-2Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilEyeClosedCircleFilled0)"/></g>`),
		g.Group(children),
	)
}