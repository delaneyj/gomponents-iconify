package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AzOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#3f9c35" d="M0 0h512v512H0z"/><path fill="#ed2939" d="M0 0h512v341.3H0z"/><path fill="#00b9e4" d="M0 0h512v170.7H0z"/><circle cx="238.8" cy="256" r="76.8" fill="#fff"/><circle cx="255.9" cy="256" r="64" fill="#ed2939"/><path fill="#fff" d="m324.2 213.3l8.1 23l22-10.5l-10.4 22l23 8.2l-23 8.2l10.4 22l-22-10.5l-8.1 23l-8.2-23l-22 10.5l10.5-22l-23-8.2l23-8.2l-10.5-22l22 10.5l8.2-23z"/>`),
		g.Group(children),
	)
}