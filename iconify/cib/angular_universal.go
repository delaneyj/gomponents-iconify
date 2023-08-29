package cib

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AngularUniversal(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M20.803 15.041v1.917a.64.64 0 0 1-.641.64h-8.323a.64.64 0 0 1-.641-.64v-1.917a.64.64 0 0 1 .641-.64h8.323a.64.64 0 0 1 .641.64zM16 20.803c-2.136 0-2.136 3.197 0 3.197s2.136-3.197 0-3.197zm4.161-11.204h-8.323a.64.64 0 0 0-.641.641v1.921a.64.64 0 0 0 .641.641h8.323a.64.64 0 0 0 .641-.641V10.24a.64.64 0 0 0-.641-.641zM30.88 5.328l-2.287 19.645L15.969 32L3.36 24.973L1.12 5.328L15.969 0zm-8.479 2.989a1.92 1.92 0 0 0-1.923-1.916h-8.957a1.92 1.92 0 0 0-1.923 1.916v15.365a1.92 1.92 0 0 0 1.923 1.916h8.957a1.92 1.92 0 0 0 1.923-1.916z"/>`),
		g.Group(children),
	)
}