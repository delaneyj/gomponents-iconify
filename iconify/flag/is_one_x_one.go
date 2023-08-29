package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IsOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagIs1x10"><path fill-opacity=".7" d="M85.4 0h486v486h-486z"/></clipPath></defs><g fill-rule="evenodd" stroke-width="0" clip-path="url(#flagIs1x10)" transform="translate(-90) scale(1.0535)"><path fill="#003897" d="M0 0h675v486H0z"/><path fill="#fff" d="M0 189h189V0h108v189h378v108H297v189H189V297H0V189z"/><path fill="#d72828" d="M0 216h216V0h54v216h405v54H270v216h-54V270H0v-54z"/></g>`),
		g.Group(children),
	)
}