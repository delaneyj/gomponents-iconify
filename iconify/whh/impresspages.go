package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Impresspages(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832.63 384q-80 0-136-56t-56-135.5t56-136t136-56.5t136 56.5t56 136t-56 135.5t-136 56zm-320 320l-192-192l192-192l192 192zm-319.5-320q-79.5 0-136-56T.63 192.5t56-136t136-56.5t136 56.5t56 136t-56 135.5t-135.5 56zm0 256q79.5 0 135.5 56.5t56 136t-56 135.5t-135.5 56t-136-56T.63 832.5t56.5-136t136-56.5zm639.5 0q80 0 136 56.5t56 136t-56 135.5t-136 56t-136-56t-56-136t56-136t136-56z"/>`),
		g.Group(children),
	)
}