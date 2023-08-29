package cryptocurrency_color

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ella(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<g fill="none" fill-rule="evenodd"><circle cx="16" cy="16" r="16" fill="#396a28"/><path fill="#fff" d="M18.073 13.212L16 11.14l-2.087 2.086l-2.164-2.164L16 5l4.28 6.005zm-4.883.738L11.14 16l2.072 2.073l-2.156 2.156L5 16l5.99-4.25zm5.584 4.137L20.86 16l-2.064-2.064l2.226-2.226L27 16l-6.044 4.269zm-4.838.71L16 20.86l2.05-2.05l2.182 2.183L16 27l-4.204-6.064zM16 12.527L19.472 16L16 19.472L12.528 16z"/></g>`),
		g.Group(children),
	)
}