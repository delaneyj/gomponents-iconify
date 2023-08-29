package emojione

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NoMobilePhones(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<circle cx="32" cy="32" r="30" fill="#ff5a79"/><circle cx="32" cy="32" r="24" fill="#333"/><path fill="#fff" d="M37.8 15H26.2c-2.1 0-3.8 1.7-3.8 3.8v26.4c0 2.1 1.7 3.8 3.8 3.8h11.5c2.1 0 3.8-1.7 3.8-3.8V18.8c.1-2.1-1.6-3.8-3.7-3.8m-7.7 1.4h3.8c.3 0 .5.2.5.5s-.2.5-.5.5h-3.8c-.3 0-.5-.2-.5-.5s.3-.5.5-.5M32 47.1c-.6 0-1.2-.5-1.2-1.1s.5-1.1 1.2-1.1c.6 0 1.2.5 1.2 1.1s-.6 1.1-1.2 1.1m7.7-4.1H24.3V18.8h15.4V43"/><path fill="#ff5a79" d="m9.23 13.474l4.243-4.243l41.295 41.295l-4.243 4.242z"/>`),
		g.Group(children),
	)
}