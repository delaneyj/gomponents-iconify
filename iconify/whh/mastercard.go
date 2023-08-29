package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Mastercard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M960.31 768h-896q-26 0-45-18.5T.31 704V64q0-26 18.5-45t45.5-19h896q27 0 45.5 19t18.5 45v640q0 26-18.5 45t-45.5 19zm-320-576q-73 0-128 50q-55-50-128-50q-80 0-136 56.5t-56 136t56 135.5t136 56q73 0 128-50q55 50 128 50q79 0 135.5-56t56.5-135.5t-56.5-136t-135.5-56.5zm0 320q-53 0-90.5-37.5t-37.5-90.5t37.5-90.5t90.5-37.5t90.5 37.5t37.5 90.5t-37.5 90.5t-90.5 37.5z"/>`),
		g.Group(children),
	)
}