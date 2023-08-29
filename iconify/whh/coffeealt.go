package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Coffeealt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M832 256L704 960q-4 24-21.5 44t-42.5 20H256q-25 0-42.5-20T192 960L64 256q-27 0-45.5-19T0 191.5t18.5-45T64 128V64q0-27 18.5-45.5T128 0h640q26 0 45 18.5T832 64v64q26 0 45 18.5t19 45t-19 45.5t-45 19z"/>`),
		g.Group(children),
	)
}