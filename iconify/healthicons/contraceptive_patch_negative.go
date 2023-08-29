package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ContraceptivePatchNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsContraceptivePatchNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM14 6a8 8 0 0 0-8 8v20a8 8 0 0 0 8 8h20a8 8 0 0 0 8-8V14a8 8 0 0 0-8-8H14Zm4.1 10.1h11.8a2 2 0 0 1 2 2v11.8a2 2 0 0 1-2 2H18.1a2 2 0 0 1-2-2V18.1a2 2 0 0 1 2-2Zm-4 2a4 4 0 0 1 4-4h11.8a4 4 0 0 1 4 4v11.8a4 4 0 0 1-4 4H18.1a4 4 0 0 1-4-4V18.1Zm7.2.5a.9.9 0 1 1-1.8 0a.9.9 0 0 1 1.8 0Zm.9 3.6a.9.9 0 1 0 0-1.8a.9.9 0 0 0 0 1.8Zm-2.7-.9a.9.9 0 1 1-1.8 0a.9.9 0 0 1 1.8 0Zm.9 3.6a.9.9 0 1 0 0-1.8a.9.9 0 0 0 0 1.8Zm2.7 1.8a.9.9 0 1 1-1.8 0a.9.9 0 0 1 1.8 0Zm-4.5.9a.9.9 0 1 0 0-1.8a.9.9 0 0 0 0 1.8Zm2.7 1.8a.9.9 0 1 1-1.8 0a.9.9 0 0 1 1.8 0Zm2.7-9.9a.9.9 0 1 0 0-1.8a.9.9 0 0 0 0 1.8Zm2.7 1.8a.9.9 0 1 1-1.8 0a.9.9 0 0 1 1.8 0ZM24 24.9a.9.9 0 1 0 0-1.8a.9.9 0 0 0 0 1.8Zm2.7 1.8a.9.9 0 1 1-1.8 0a.9.9 0 0 1 1.8 0ZM24 30.3a.9.9 0 1 0 0-1.8a.9.9 0 0 0 0 1.8Zm4.5-11.7a.9.9 0 1 1-1.8 0a.9.9 0 0 1 1.8 0Zm.9 3.6a.9.9 0 1 0 0-1.8a.9.9 0 0 0 0 1.8Zm-.9 1.8a.9.9 0 1 1-1.8 0a.9.9 0 0 1 1.8 0Zm.9 3.6a.9.9 0 1 0 0-1.8a.9.9 0 0 0 0 1.8Zm-.9 1.8a.9.9 0 1 1-1.8 0a.9.9 0 0 1 1.8 0Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsContraceptivePatchNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}