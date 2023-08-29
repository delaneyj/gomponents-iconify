package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PenCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilPenCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M6.944 14.79a.5.5 0 0 1 .141-.277L17.163 4.435a.5.5 0 0 1 .707 0l3.89 3.89a.5.5 0 0 1 0 .706L11.68 19.11a.5.5 0 0 1-.277.14l-4.595.706a.5.5 0 0 1-.57-.57l.705-4.594Zm.964.314l-.577 3.76l3.759-.578l9.608-9.608l-3.181-3.182l-9.61 9.608Z"/><path d="m18.472 11.173l-3.537-3.53l.707-.708l3.536 3.53l-.706.708Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilPenCircleFilled0)"/></g>`),
		g.Group(children),
	)
}