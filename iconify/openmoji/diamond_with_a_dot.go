package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DiamondWithADot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#92d3f5" d="m12.065 36.5l24.218-24.218L60.5 36.5L36.283 60.718z"/><path fill="#e27022" d="M31.757 32.757h8.485v8.485h-8.485z"/><g fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path d="m32 33l-7.5-7.5M40 33l8.5-8.5M40 41l7.5 7.5M32 41l-7.5 7.5m-12.435-12l24.218-24.218L60.5 36.5L36.283 60.718z"/><path d="M31.757 32.757h8.485v8.485h-8.485z"/></g>`),
		g.Group(children),
	)
}