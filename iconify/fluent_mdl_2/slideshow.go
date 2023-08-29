package fluent_mdl_2

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Slideshow(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 2048 2048"),
		g.Raw(`<path fill="currentColor" d="m1280 1024l-512 256V768l512 256zm-416 101l201-101l-201-101v202zm1184-869v1536H0V256h2048zm-128 128H128v1280h1792V384zm-128 1152H256V512h1536v1024zm-128-896H384v768h1280V640z"/>`),
		g.Group(children),
	)
}