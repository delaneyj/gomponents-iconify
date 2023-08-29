package whh

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Copyapp(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 1024 1024"),
		g.Raw(`<path fill="currentColor" d="m832 448l192 384q-179-12-329-38q-57-10-118-27l-129 65l-30-76l-66 44l-161-437l-173 58l-18-37l145-145l-17-47l74 25l-10-25l96 32l118 245q51-38 73-50l65-35q26 8 57 18L896 0L796 449z"/>`),
		g.Group(children),
	)
}