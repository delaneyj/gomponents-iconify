package fe

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RiceCracker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g id="feRiceCracker0" fill="none" fill-rule="evenodd" stroke="none" stroke-width="1"><g id="feRiceCracker1" fill="currentColor" fill-rule="nonzero"><path id="feRiceCracker2" d="M12 22C6.477 22 2 17.523 2 12S6.477 2 12 2s10 4.477 10 10s-4.477 10-10 10Zm-5-3.755V13a2 2 0 0 1 2-2h6a2 2 0 0 1 2 2v5.245a8 8 0 1 0-10 0Z"/></g></g>`),
		g.Group(children),
	)
}