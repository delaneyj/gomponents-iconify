package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Fonttypewriter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832.053 384q-26.5 0-45-18.5t-18.5-45.5V128h-256v768h128q26 0 45 19t19 45.5t-19 45t-45 18.5h-384q-27 0-45.5-18.5t-18.5-45t18.5-45.5t45.5-19h128V128h-256v192q0 27-19 45.5t-45.5 18.5t-45-18.5T.553 320V64q0-26 18.5-45t45.5-19h768q23 0 23 5q19 7 30 23t11 36v256q0 27-19 45.5t-45.5 18.5z"/>`),
		g.Group(children),
	)
}