package fxemoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Emptypages(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="#EAC083" d="M446.067 512h-217.97c-2.818 0-5.102-3.663-5.102-8.182V8.182c0-4.519 2.284-8.182 5.102-8.182H446.23c2.727 0 4.938 3.546 4.938 7.92v495.898c.001 4.519-2.283 8.182-5.101 8.182z"/><path fill="#F9E7C0" d="M392.753 512H191.967L35.007 355.04V8.182A8.182 8.182 0 0 1 43.189 0h349.826a7.92 7.92 0 0 1 7.92 7.92v495.898a8.182 8.182 0 0 1-8.182 8.182z"/><path fill="#EAC083" d="M191.967 512L35.007 355.04h123.597c18.426 0 33.363 14.937 33.363 33.363V512z"/>`),
		g.Group(children),
	)
}