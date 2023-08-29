package fluent

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DocumentFooterRemoveTwentyFourRegular(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 20 20"),
		g.Raw(`<g fill="none"><path d="M7 15.75c0-.966.784-1.75 1.75-1.75h6.5c.967 0 1.75.784 1.75 1.75v1.5A1.75 1.75 0 0 1 15.25 19h-6.5A1.75 1.75 0 0 1 7 17.25v-1.5zm1.75-.25a.25.25 0 0 0-.25.25v1.5c0 .138.112.25.25.25h6.5a.25.25 0 0 0 .25-.25v-1.5a.25.25 0 0 0-.25-.25h-6.5z" fill="currentColor"/><path d="M17.5 1a5.5 5.5 0 1 1 0 11a5.5 5.5 0 0 1 0-11zm-2.476 3.024a.5.5 0 0 0 0 .707l1.77 1.77l-1.767 1.766a.5.5 0 1 0 .707.708L17.5 7.208l1.77 1.769a.5.5 0 1 0 .707-.707L18.208 6.5l1.772-1.77a.5.5 0 0 0-.707-.707L17.5 5.794l-1.77-1.77a.5.5 0 0 0-.707 0z" fill="currentColor"/><path d="M18.495 12.924a6.459 6.459 0 0 0 1.5-.42v7.24a2.25 2.25 0 0 1-2.096 2.245l-.154.005h-11.5A2.25 2.25 0 0 1 4 19.898l-.005-.154V4.246A2.25 2.25 0 0 1 6.091 2l.154-.005h6.569a6.52 6.52 0 0 0-1.08 1.5H6.246a.75.75 0 0 0-.743.648l-.007.102v15.498c0 .38.282.693.648.743l.102.007h11.5a.75.75 0 0 0 .743-.648l.007-.102v-6.82z" fill="currentColor"/></g>`),
		g.Group(children),
	)
}