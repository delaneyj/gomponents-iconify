package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wallet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M896.56 895h-576q-53 0-90.5-37.5t-37.5-90.5h-64q-53 0-90.5-37.5T.56 639V127q0-53 37.5-90t90.5-37h768q53 0 90.5 37t37.5 90v640q0 53-37.5 90.5t-90.5 37.5zm-480-256q96 0 96-96t-96-96t-96 96t96 96zm448-512h-647q-37 0-63 25t-26 61v341q0 29 18 51.5t46 29.5V319q0-53 37.5-90.5t90.5-37.5h544q13 0 22.5-9t9.5-22.5t-9.5-23t-22.5-9.5z"/>`),
		g.Group(children),
	)
}