package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Video(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m2048 1544l-512-256v248H0V512h1536v248l512-256v1040zm-640-904H128v768h1280V640zm512 71l-384 193v240l384 193V711z"/>`),
		g.Group(children),
	)
}