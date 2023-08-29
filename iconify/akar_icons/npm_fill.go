package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func NpmFill(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><g clip-path="url(#akarIconsNpmFill0)"><path fill="currentColor" fill-rule="evenodd" d="M24 0H0v24h24V0ZM2.578 2.578H21.42V21.42h-4.75V7.33h-4.752v14.09h-9.34V2.578Z" clip-rule="evenodd"/></g><defs><clipPath id="akarIconsNpmFill0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}