package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BookCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilBookCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M13.08 7.304L5.244 6.019A1.5 1.5 0 0 0 3.5 7.5v9.738a1.5 1.5 0 0 0 1.268 1.482l8.155 1.275a.5.5 0 0 0 .577-.494V7.797a.5.5 0 0 0-.42-.493Zm-8-.298l7.42 1.216v10.694L4.923 17.73a.5.5 0 0 1-.423-.493V7.5a.5.5 0 0 1 .58-.494Z"/><path d="M21 6a1.5 1.5 0 0 0-.243.02L12.92 7.303a.5.5 0 0 0-.419.493V19.5a.5.5 0 0 0 .577.494l8.155-1.276a1.5 1.5 0 0 0 1.268-1.481V7.5A1.5 1.5 0 0 0 21 6Zm.077 11.73L13.5 18.916V8.222l7.42-1.216a.501.501 0 0 1 .58.494v9.737a.5.5 0 0 1-.423.494Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilBookCircleFilled0)"/></g>`),
		g.Group(children),
	)
}