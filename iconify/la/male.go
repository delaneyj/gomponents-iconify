package la

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Male(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M16 2c-2.2 0-4 1.8-4 4a3.96 3.96 0 0 0 1.125 2.75C11.273 9.773 10 11.746 10 14v5.406l.281.313L12 21.437V30h2v-9.406l-.281-.313L12 18.563V14c0-2.219 1.781-4 4-4c2.219 0 4 1.781 4 4v4.563l-1.719 1.718l-.281.313V30h2v-8.563l1.719-1.718l.281-.313V14c0-2.254-1.273-4.227-3.125-5.25A3.958 3.958 0 0 0 20 6c0-2.2-1.8-4-4-4zm0 2c1.117 0 2 .883 2 2s-.883 2-2 2s-2-.883-2-2s.883-2 2-2z"/>`),
		g.Group(children),
	)
}