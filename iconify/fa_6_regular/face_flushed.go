package fa_6_regular

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FaceFlushed(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M464 256a208 208 0 1 1-416 0a208 208 0 1 1 416 0zM256 0a256 256 0 1 0 0 512a256 256 0 1 0 0-512zm-95.6 248a24 24 0 1 0 0-48a24 24 0 1 0 0 48zm216-24a24 24 0 1 0-48 0a24 24 0 1 0 48 0zM192 336c-13.3 0-24 10.7-24 24s10.7 24 24 24h128c13.3 0 24-10.7 24-24s-10.7-24-24-24H192zm-32-160a48 48 0 1 1 0 96a48 48 0 1 1 0-96zm0 128a80 80 0 1 0 0-160a80 80 0 1 0 0 160zm144-80a48 48 0 1 1 96 0a48 48 0 1 1-96 0zm128 0a80 80 0 1 0-160 0a80 80 0 1 0 160 0z"/>`),
		g.Group(children),
	)
}