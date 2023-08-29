package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SmartphoneHomeButtonCircleFilled(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="none"><defs><mask id="pepiconsPopSmartphoneHomeButtonCircleFilled0"><path fill="#fff" d="M0 0h26v26H0z"/><g fill="#000"><path fill-rule="evenodd" d="M17.5 3h-9a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h9a2 2 0 0 0 2-2V5a2 2 0 0 0-2-2Zm-9 18V5h9v16h-9Z" clip-rule="evenodd"/><path d="M13 20.25a1.25 1.25 0 1 1 0-2.5a1.25 1.25 0 0 1 0 2.5Z"/></g></mask></defs><circle cx="13" cy="13" r="13" fill="currentColor" mask="url(#pepiconsPopSmartphoneHomeButtonCircleFilled0)"/></g>`),
		g.Group(children),
	)
}