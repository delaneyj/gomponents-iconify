package akar_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Lifesaver(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none"><g stroke="currentColor" stroke-linejoin="round" stroke-width="2" clip-path="url(#akarIconsLifesaver0)"><circle cx="12" cy="12" r="10" stroke-linecap="round" transform="rotate(45 12 12)"/><circle cx="12" cy="12" r="4" stroke-linecap="round" transform="rotate(45 12 12)"/><path d="m19.071 4.929l-4.243 4.243m-5.656 5.656l-4.243 4.243m14.142 0l-4.243-4.243M9.172 9.172L4.929 4.929"/></g><defs><clipPath id="akarIconsLifesaver0"><path fill="#fff" d="M0 0h24v24H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}