package pepicons_pencil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func YenCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPencilYenCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000" fill-rule="evenodd" clip-rule="evenodd"><path d="M12.336 11.87a.5.5 0 0 1-.706-.034l-5-5.512a.5.5 0 1 1 .74-.672l5 5.512a.5.5 0 0 1-.034.706Z"/><path d="M11.664 11.87a.5.5 0 0 0 .706-.034l5-5.512a.5.5 0 1 0-.74-.672l-5 5.512a.5.5 0 0 0 .034.706Z"/><path d="M6.5 11.988a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H7a.5.5 0 0 1-.5-.5Zm0 3.5a.5.5 0 0 1 .5-.5h10a.5.5 0 0 1 0 1H7a.5.5 0 0 1-.5-.5Z"/><path d="M12 11a.5.5 0 0 1 .5.5v8.488a.5.5 0 0 1-1 0V11.5a.5.5 0 0 1 .5-.5Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPencilYenCircleFilled0)"/></g>`),
		g.Group(children),
	)
}