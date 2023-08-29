package radix_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LetterCaseCapitalize(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 15 15"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M3.69 2.75a.5.5 0 0 1 .467.32l3.21 8.32a.5.5 0 0 1-.933.36L5.383 9.025H2.01L.967 11.75a.5.5 0 0 1-.934-.358l3.19-8.32a.5.5 0 0 1 .467-.321Zm.002 1.893l1.363 3.532H2.337l1.355-3.532Zm7.207.564c-1.64 0-2.89 1.479-2.89 3.403c0 2.024 1.35 3.402 2.89 3.402a3.06 3.06 0 0 0 2.255-.99v.508a.45.45 0 1 0 .9 0V5.72a.45.45 0 0 0-.9 0v.503A3.01 3.01 0 0 0 10.9 5.207Zm2.255 4.591V7.302c-.39-.721-1.213-1.244-2.067-1.244c-.978 0-2.052.908-2.052 2.552c0 1.543.974 2.552 2.052 2.552c.883 0 1.685-.667 2.067-1.364Z" clip-rule="evenodd"/>`),
		g.Group(children),
	)
}