package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WashMachineTwo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M16.031.042h-15v3h15v-3zM3 1.958H2v-1h1v1zm2 0H4v-1h1v1z"/><path d="M1.021 3v13h14.958V3H1.021zm7.436 12a5.5 5.5 0 1 1 0-11a5.5 5.5 0 0 1 0 11z"/><path d="M8.516 5.083c-2.453 0-4.441 1.979-4.441 4.416c0 2.438 1.988 4.418 4.441 4.418c2.454 0 4.442-1.979 4.442-4.418c0-2.437-1.988-4.416-4.442-4.416zm3.283 4.945c.079 0 .15-.021.218-.048c-.235 1.716-2.19 1.256-3.501.037c-1.531-1.424-3.537 1.425-3.537-.519c0-1.94 1.584-3.516 3.537-3.516a3.527 3.527 0 0 1 3.501 3.036a.53.53 0 1 0-.218 1.01z"/></g>`),
		g.Group(children),
	)
}