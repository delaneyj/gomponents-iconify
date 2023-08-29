package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Imo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 2.5C12.126 2.5 2.5 12.126 2.5 24c0 5.49 2.075 10.484 5.462 14.283c-.971 2.801-2.994 3.675-5.462 4.567c2.26 2.182 6.668 2.76 10.623-.327A21.368 21.368 0 0 0 24 45.5c11.874 0 21.5-9.626 21.5-21.5S35.874 2.5 24 2.5Z"/><rect width="5.818" height="7.709" x="31.64" y="19.888" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.909" ry="2.909"/><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M16.918 22.797a2.909 2.909 0 0 1 2.909-2.91h0a2.909 2.909 0 0 1 2.909 2.91v4.8m-5.818-7.709v7.709"/><path d="M22.736 22.797a2.909 2.909 0 0 1 2.909-2.91h0a2.909 2.909 0 0 1 2.909 2.91v4.8"/></g><circle cx="13.56" cy="17.179" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.56 19.888v7.709"/>`),
		g.Group(children),
	)
}