package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tampermonkey(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.428 1024h-768q-53 0-90.5-37.5T.428 896V128q0-53 37.5-90.5t90.5-37.5h768q53 0 90.5 37.5t37.5 90.5v768q0 53-37.5 90.5t-90.5 37.5zm-640-448q-80 0-136 56t-56 136t56 136t136 56t136-56t56-136t-56-136t-136-56zm512 0q-80 0-136 56t-56 136t56 136t136 56t136-56t56-136t-56-136t-136-56z"/>`),
		g.Group(children),
	)
}