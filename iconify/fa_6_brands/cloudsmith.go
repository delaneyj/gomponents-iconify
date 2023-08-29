package fa_6_brands

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloudsmith(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 448 512"),
		g.Raw(`<path fill="currentColor" d="M512 227.6v56.9L284.4 512h-56.8L0 284.4v-56.8L227.6 0h56.9L512 227.6zm-256 162a133.6 133.6 0 1 0 0-267.1a133.6 133.6 0 1 0 0 267.1z"/>`),
		g.Group(children),
	)
}