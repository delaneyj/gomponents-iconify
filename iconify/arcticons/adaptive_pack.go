package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdaptivePack(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.5 2.98C11.24 4.74 4.74 11.24 2.98 19.5m42.04 0C43.26 11.24 36.76 4.74 28.5 2.98m0 42.04c8.26-1.76 14.76-8.26 16.52-16.52m-42.04 0c1.76 8.26 8.26 14.76 16.52 16.52"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M18.394 29.096a1.959 1.959 0 1 1 1.943-1.879a1.943 1.943 0 0 1-1.943 1.88Zm11.339 0a1.959 1.959 0 1 1 1.943-1.879a1.943 1.943 0 0 1-1.943 1.88Z"/><path d="M24.032 17.104h0c7.67 0 13.887 6.218 13.887 13.887h0c0 1.09-.884 1.975-1.975 1.975H12.056a1.975 1.975 0 0 1-1.975-1.975h0c0-7.67 6.217-13.887 13.887-13.887h.064Zm-11.976-2.07l3.949 4.634m20.035-4.634l-3.934 4.634"/></g>`),
		g.Group(children),
	)
}