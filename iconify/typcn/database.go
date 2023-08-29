package typcn

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Database(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M18.983 8.791C18.807 5.087 15.747 2.125 12 2.125S5.193 5.087 5.017 8.791L5 8.875v6.25c0 3.86 3.141 7 7 7s7-3.14 7-7v-6.25l-.017-.084zM12 17.625a5.502 5.502 0 0 1-5-3.222v-.388a6.979 6.979 0 0 0 10 0v.388a5.502 5.502 0 0 1-5 3.222zm0-13.5c2.757 0 5 2.243 5 5s-2.243 5-5 5s-5-2.243-5-5s2.243-5 5-5zm0 16c-2.271 0-4.172-1.532-4.778-3.609C8.41 17.809 10.11 18.625 12 18.625s3.59-.816 4.778-2.109c-.606 2.077-2.507 3.609-4.778 3.609z"/>`),
		g.Group(children),
	)
}