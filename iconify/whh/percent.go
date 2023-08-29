package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Percent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832.02 1024q-80 0-136-56.5t-56-136t56-135.5t136-56t136 56t56 135.5t-56 136t-136 56.5zm.5-256q-26.5 0-45.5 18.5t-19 45t19 45.5t45.5 19t45-19t18.5-45.5t-18.5-45t-45-18.5zm-436.5 192q-14 27-43.5 45.5t-56.5 18.5h-129q-27 0-35.5-19t5.5-45l491-896q14-27 43.5-45.5T728.02 0h129q27 0 35.5 18.5t-5.5 45.5zm-204-576q-80 0-136-56.5t-56-136t56-135.5t136-56t136 56t56 135.5t-56 136t-136 56.5zm-.5-256q-26.5 0-45 18.5t-18.5 45t18.5 45.5t45.5 19t45.5-19t18.5-45.5t-19-45t-45.5-18.5z"/>`),
		g.Group(children),
	)
}