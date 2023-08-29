package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CircledZeroWithSlash(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<circle cx="36" cy="36" r="26.68" fill="#fff" fill-rule="evenodd"/><g fill="none" stroke="#000" stroke-linejoin="round"><circle cx="36" cy="36" r="26.68" stroke-linecap="round" stroke-width="4.74"/><path stroke-width="4.74" d="m43.24 22.41l-15.01 26l-.327-.189"/><path stroke-linecap="round" stroke-miterlimit="10" stroke-width="7.998" d="M36 51.209a8.327 8.327 0 0 1-8.329-8.329V29.11a8.33 8.33 0 0 1 16.658 0v13.77A8.327 8.327 0 0 1 36 51.21z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}