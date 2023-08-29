package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DatabaseUpload(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><g transform="translate(2 1)"><ellipse cx="6.43" cy="2.421" rx="6.43" ry="2.421"/><path d="M6.479 5.988C2.963 5.988.032 4.986.032 4.314V7.66c0 1.214 2.887 2.196 6.447 2.196s6.447-.982 6.447-2.196V4.314c0 .672-2.932 1.674-6.447 1.674z"/></g><path d="M11.953 15.031H9.906l2.705-3.488c-1.131.276-2.587.475-4.164.475c-3.516 0-6.447-.98-6.447-1.636v3.269c0 1.188 2.887 2.146 6.447 2.146c1.32 0 2.484-.133 3.506-.359v-.407z"/><path d="M13.018 15.969v-1.988h-1.002l1.492-1.965l1.476 1.965h-1.017v1.988h-.949z"/></g>`),
		g.Group(children),
	)
}