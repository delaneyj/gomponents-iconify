package simple_icons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Tamiya(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M0 6.408v4.27h4.496l1.36-4.27Zm5.856 0l1.398 4.27h4.496v-4.27Zm5.894 4.27l-3.627 2.644l1.398 4.27h2.23Zm-2.23 6.914l-3.664-2.645l-3.627 2.645Zm-7.291 0l1.398-4.27L0 10.678v6.914zM12.25 6.408v4.27h4.496l1.36-4.27zm5.856 0l1.398 4.27H24v-4.27ZM24 10.678l-3.627 2.644l1.398 4.27H24Zm-2.23 6.914l-3.664-2.645l-3.627 2.645zm-7.29 0l1.397-4.27l-3.627-2.644v6.914z"/>`),
		g.Group(children),
	)
}