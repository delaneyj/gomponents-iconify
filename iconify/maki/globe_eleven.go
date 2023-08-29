package maki

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GlobeEleven(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path d="M6.5 10H6V8.95a4.478 4.478 0 0 0 2.682-1.268L8.15 7.15a3.739 3.739 0 0 1-2.65 1.1A3.754 3.754 0 0 1 1.75 4.5c0-1.034.42-1.971 1.1-2.65l-.532-.532A4.486 4.486 0 0 0 1 4.5c0 2.314 1.753 4.198 4 4.45V10h-.5a.5.5 0 1 0 0 1h2a.5.5 0 0 0 0-1z" fill="currentColor"/><path d="M5.5 7a2.5 2.5 0 1 0 0-5a2.5 2.5 0 0 0 0 5zm-.545-4.415l.493.616v.837l.613 1.111h1.322c-.17.492-.522.892-.98 1.126L5.5 5.5h-1l-.968-1.32a1.991 1.991 0 0 1 1.423-1.595z" fill="currentColor"/>`),
		g.Group(children),
	)
}