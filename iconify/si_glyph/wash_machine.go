package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WashMachine(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M16.031.042h-15v3h15v-3zm-12.989 2H1.938V.958h1.104v1.084zm2 0H3.938V.958h1.104v1.084z"/><path d="M1.021 3v13h14.958V3H1.021zm7.491 12.083c-3.068 0-5.555-2.5-5.555-5.583c0-3.084 2.486-5.584 5.555-5.584c3.067 0 5.554 2.5 5.554 5.584c-.001 3.083-2.487 5.583-5.554 5.583z"/><path d="M8.516 5.083c-2.453 0-4.441 1.979-4.441 4.416c0 2.438 1.988 4.418 4.441 4.418c2.454 0 4.442-1.979 4.442-4.418c0-2.437-1.988-4.416-4.442-4.416zm3.283 4.945c.079 0 .15-.021.218-.048a3.527 3.527 0 0 1-3.501 3.037a3.528 3.528 0 0 1-3.537-3.519c0-1.94 1.584-3.516 3.537-3.516a3.527 3.527 0 0 1 3.501 3.036a.53.53 0 1 0-.218 1.01z"/></g>`),
		g.Group(children),
	)
}