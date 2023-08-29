package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CodeCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilCodeCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M4.962 12.666a.5.5 0 0 1 .706-.038l3.333 3a.5.5 0 1 1-.669.744l-3.333-3a.5.5 0 0 1-.037-.706Z"/><path d="M9.038 9.666a.5.5 0 0 1-.037.706l-3.333 3a.5.5 0 0 1-.67-.744l3.334-3a.5.5 0 0 1 .706.038Zm12 3a.5.5 0 0 1-.037.706l-3.333 3a.5.5 0 0 1-.67-.744l3.334-3a.5.5 0 0 1 .706.038Z"/><path d="M16.962 9.666a.5.5 0 0 1 .706-.038l3.333 3a.5.5 0 0 1-.669.744l-3.333-3a.5.5 0 0 1-.037-.706Zm-2.33-2.648a.5.5 0 0 1 .35.614l-3 11a.5.5 0 0 1-.964-.264l3-11a.5.5 0 0 1 .614-.35Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilCodeCircleFilled0)"/></g>`),
		g.Group(children),
	)
}