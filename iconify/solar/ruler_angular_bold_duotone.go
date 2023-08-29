package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RulerAngularBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="currentColor"><path d="M6 2h12.75c1.24.001 1.888.02 2.361.337a2 2 0 0 1 .552.552C22 3.393 22 4.096 22 5.5c0 1.404 0 2.107-.337 2.611a2 2 0 0 1-.552.552C20.607 9 19.904 9 18.5 9H11c-.943 0-1.414 0-1.707.293C9 9.586 9 10.057 9 11v7.5c0 1.404 0 2.107-.337 2.611a2 2 0 0 1-.552.552C7.607 22 6.904 22 5.5 22c-1.404 0-2.107 0-2.611-.337a2 2 0 0 1-.552-.552c-.316-.473-.336-1.121-.337-2.361V6c0-1.886 0-2.828.586-3.414C3.172 2 4.114 2 6 2Z" opacity=".5"/><path d="M4 17.25H2v1.5h2a.75.75 0 0 0 0-1.5Zm1-3H2v1.5h3a.75.75 0 0 0 0-1.5Zm-1-3H2v1.5h2a.75.75 0 0 0 0-1.5Zm1-3H2v1.5h3a.75.75 0 0 0 0-1.5ZM9.75 5V2h-1.5v3a.75.75 0 0 0 1.5 0Zm3-1V2h-1.5v2a.75.75 0 0 0 1.5 0Zm3 1V2h-1.5v3a.75.75 0 0 0 1.5 0Zm3-1V2h-1.5v2a.75.75 0 0 0 1.5 0Z"/></g>`),
		g.Group(children),
	)
}