package flat_ui

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PaperBag(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 100 100"),
		g.Raw(`<path fill="#F39C12" d="M100 93a6 6 0 0 1-6 6H6a6 6 0 0 1-6-6V20h100v73z"/><path fill="#E57D22" d="M4 0h92v20H4z"/><path fill="#F39C12" d="M96 20h-8V8l8-8zM4 20h8V8L4 0z"/><path fill="#F1C40F" d="M12 8v12H0zm76 0v12h12z"/><circle cx="26" cy="40" r="5" opacity=".3"/><circle cx="74" cy="40" r="5" opacity=".3"/><path fill="none" stroke="#fff" stroke-linecap="round" stroke-miterlimit="10" stroke-width="6" d="M74 40c0 13.254-10.745 24-24 24S26 53.255 26 40" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}