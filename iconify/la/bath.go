package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Bath(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M26 4c-2.21 0-4 1.79-4 4h-2v2h6V8h-2c0-1.191.809-2 2-2c1.191 0 2 .809 2 2v6H1v2h1.188l1.53 7.594v.031a3.062 3.062 0 0 0 2 2.219L5 28h2l.656-2h16.688L25 28h2l-.719-2.156c1.047-.32 1.86-1.16 2.094-2.219v-.031L29.813 16H31v-2h-1V8c0-2.21-1.79-4-4-4zM4.219 16h23.593l-1.406 7.219c-.117.433-.484.781-1 .781H6.688c-.536 0-.899-.355-1-.813z"/>`),
		g.Group(children),
	)
}