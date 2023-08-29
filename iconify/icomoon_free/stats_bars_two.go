package icomoon_free

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func StatsBarsTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 0 0"),
		g.Raw(`<path fill="currentColor" d="M4.5 6h-3c-.275 0-.5.225-.5.5v9c0 .275.225.5.5.5h3c.275 0 .5-.225.5-.5v-9c0-.275-.225-.5-.5-.5zm0 9h-3v-4h3v4zm5-11h-3c-.275 0-.5.225-.5.5v11c0 .275.225.5.5.5h3c.275 0 .5-.225.5-.5v-11c0-.275-.225-.5-.5-.5zm0 11h-3v-5h3v5zm5-13h-3c-.275 0-.5.225-.5.5v13c0 .275.225.5.5.5h3c.275 0 .5-.225.5-.5v-13c0-.275-.225-.5-.5-.5zm0 13h-3V9h3v6z"/>`),
		g.Group(children),
	)
}