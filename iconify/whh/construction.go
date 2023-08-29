package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Construction(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="M1024.585 64v192q0 27-18.5 45.5t-45.5 18.5h-64v480q0 13-9.5 22.5t-22.5 9.5h-64q-13 0-22.5-9.5t-9.5-22.5V576h-512v224q0 13-9.5 22.5t-22.5 9.5h-64q-13 0-22.5-9.5t-9.5-22.5V320h-64q-27 0-45.5-18.5T.585 256V64q0-27 18.5-45.5T64.585 0h896q27 0 45.5 18.5t18.5 45.5zm-768 448h512V320h-512v192zm-192-448v64l64-64h-64zm0 192h128l192-192h-128zm448-192l-192 192h128l192-192h-128zm256 0l-192 192h128l192-192h-128zm192 64l-128 128h128V128z"/>`),
		g.Group(children),
	)
}