package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SgFourXThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagSg4x30"><path fill-opacity=".7" d="M0 0h640v480H0z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagSg4x30)"><path fill="#fff" d="M-20 0h720v480H-20z"/><path fill="#df0000" d="M-20 0h720v240H-20z"/><path fill="#fff" d="M146 40.2a84.4 84.4 0 0 0 .8 165.2a86 86 0 0 1-106.6-59a86 86 0 0 1 59-106c16-4.6 30.8-4.7 46.9-.2z"/><path fill="#fff" d="m133 110l4.9 15l-13-9.2l-12.8 9.4l4.7-15.2l-12.8-9.3l15.9-.2l5-15l5 15h15.8zm17.5 52l5 15.1l-13-9.2l-12.9 9.3l4.8-15.1l-12.8-9.4l15.9-.1l4.9-15.1l5 15h16zm58.5-.4l4.9 15.2l-13-9.3l-12.8 9.3l4.7-15.1l-12.8-9.3l15.9-.2l5-15l5 15h15.8zm17.4-51.6l4.9 15.1l-13-9.2l-12.8 9.3l4.8-15.1l-12.9-9.4l16-.1l4.8-15.1l5 15h16zm-46.3-34.3l5 15.2l-13-9.3l-12.9 9.4l4.8-15.2l-12.8-9.4l15.8-.1l5-15.1l5 15h16z"/></g>`),
		g.Group(children),
	)
}