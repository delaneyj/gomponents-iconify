package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ImoBeta(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.327 33.43A21.413 21.413 0 0 0 45.5 24c0-11.874-9.626-21.5-21.5-21.5S2.5 12.126 2.5 24c0 5.49 2.075 10.484 5.462 14.283c-.971 2.801-2.994 3.675-5.462 4.567c2.26 2.182 6.668 2.76 10.623-.327A21.368 21.368 0 0 0 24 45.5c3.385 0 6.587-.782 9.435-2.176"/><rect width="5.818" height="7.709" x="31.64" y="19.888" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.909" ry="2.909"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M16.918 22.797a2.909 2.909 0 0 1 2.909-2.91h0a2.909 2.909 0 0 1 2.909 2.91v4.8m-5.818-7.709v7.709"/><path d="M22.736 22.797a2.909 2.909 0 0 1 2.909-2.91h0a2.909 2.909 0 0 1 2.909 2.91v4.8"/></g><circle cx="13.56" cy="17.179" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.56 19.888v7.709"/><circle cx="38.5" cy="38.5" r="7" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.536 38.5a2 2 0 1 1 0 4h-3.3v-8h3.3a2 2 0 1 1 0 4h0Zm0 0h-3.3"/>`),
		g.Group(children),
	)
}