package icon_park_outline

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func HeavyMetal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-width="4"><path stroke-linejoin="round" d="m22.219 7.378l11.668 9.244L36.177 24l-10.915 2.085L12.59 14.531l2.692-5.53l6.938-1.623Z" clip-rule="evenodd"/><path stroke-linecap="round" d="m15.28 9.001l11.206 9.6"/><path stroke-linecap="round" stroke-linejoin="round" d="m25.263 26.085l1.224-7.953l7.4-1.51m-8.894 14.401l2.286 7.08l-11.678 2.276L4 29.014l2.57-6.389l5.458-1.271"/><path stroke-linecap="round" d="m6.57 22.625l10.714 10.133"/><path stroke-linecap="round" stroke-linejoin="round" d="m15.601 39.865l1.885-7.733l7.505-1.088"/><path stroke-linejoin="round" d="M34.887 29.608L34 36.8l9.236-1.801l-1.955-6.812l-6.394 1.42Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}