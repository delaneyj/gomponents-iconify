package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CgOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagCg1x10"><path fill-opacity=".7" d="M115.7 0h496.1v496h-496z"/></clipPath></defs><g fill-rule="evenodd" stroke-width="1pt" clip-path="url(#flagCg1x10)" transform="translate(-119.5) scale(1.032)"><path fill="#ff0" d="M0 0h744v496H0z"/><path fill="#00ca00" d="M0 0v496L496 0H0z"/><path fill="red" d="M248 496h496V0L248 496z"/></g>`),
		g.Group(children),
	)
}