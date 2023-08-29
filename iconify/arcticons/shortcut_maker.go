package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ShortcutMaker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.5 23.43l-2.978-2.978m-1.836-1.818v-.017l-3.439-3.438v2.851m0 2.184a10.493 10.493 0 0 0-6.839 2.497m-1.97 1.951a21.171 21.171 0 0 0-3.938 8.159c-.039-.019 2.603-5.746 10.493-6.192M34.5 23.43l-2.979 2.978m-1.835 1.817v.017l-3.439 3.438v-5.058"/><circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}