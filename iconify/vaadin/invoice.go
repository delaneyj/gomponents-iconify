package vaadin

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Invoice(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4.4 10.2c-.6.1-1.4-.3-1.7-.4l-.5.9s.9.4 1.7.5v.8h1v-.9c.9-.3 1.4-1.1 1.5-1.8c0-.8-.6-1.4-1.9-1.9c-.4-.2-1.1-.5-1.1-.9c0-.5.4-.8 1-.8c.7 0 1.4.3 1.4.3l.4-.9s-.5-.2-1.2-.4V4H4v.7c-.9.2-1.5.8-1.6 1.7c0 1.2 1.3 1.7 1.8 1.9c.6.2 1.3.6 1.3.9c0 .4-.4.9-1.1 1z"/><path fill="currentColor" d="M0 2v12h16V2H0zm15 11H1V3h14v10z"/><path fill="currentColor" d="M8 5h6v1H8V5zm0 2h6v1H8V7zm0 2h3v1H8V9z"/>`),
		g.Group(children),
	)
}