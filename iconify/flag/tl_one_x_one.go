package flag

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TlOneXOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<defs><clipPath id="flagTl1x10"><path fill-opacity=".7" d="M0 0h496v496H0z"/></clipPath></defs><g fill-rule="evenodd" clip-path="url(#flagTl1x10)" transform="scale(1.0321)"><path fill="#cb000f" d="M0 0h999v496H0z"/><path fill="#f8c00c" d="M0 0c3.1 0 496 248.7 496 248.7L0 496.1V0z"/><path d="M0 0c2 0 330 248.7 330 248.7L0 496.1V0z"/><path fill="#fff" d="m181.9 288.9l-59-13L93 327l-5-57.8l-58.8-13l53.1-24l-3.2-57.5l39 42l53.6-24.4l-28 52.2l38 44.4z"/></g>`),
		g.Group(children),
	)
}