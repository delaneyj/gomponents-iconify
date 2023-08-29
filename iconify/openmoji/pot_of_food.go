package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PotOfFood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#FFF" d="M7 36c0 16.017 12.983 29 29 29s29-12.983 29-29H7z"/><path fill="#d0cfce" d="M49 36c0 16.016-4 29-13 29c16.017 0 29-12.984 29-29H49z"/><path fill="none" stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-miterlimit="10" stroke-width="2" d="M7 36c0 16.017 12.983 29 29 29s29-12.983 29-29H7zm0-10.859l7 7M29.043 7.846a4.953 4.953 0 0 0 1.42 3.084a5 5 0 0 1 0 7.07a5 5 0 0 0 0 7.07a5 5 0 0 1 0 7.07M34.57 7.036c-.1 1.398.362 2.826 1.43 3.894A5 5 0 0 1 36 18a5 5 0 0 0 0 7.07a5 5 0 0 1 0 7.07m4.01-24.857a4.969 4.969 0 0 0 1.454 3.647a5 5 0 0 1 0 7.07a5 5 0 0 0 0 7.07a5 5 0 0 1 0 7.07"/>`),
		g.Group(children),
	)
}