package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Newspaper(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M88 72v316a20 20 0 0 1-40 0V152H16v236a52.059 52.059 0 0 0 52 52h376a52.059 52.059 0 0 0 52-52V72Zm376 316a20.023 20.023 0 0 1-20 20H116a51.722 51.722 0 0 0 4-20V104h344Z"/><path fill="currentColor" d="M296 136H152v160h144Zm-32 128h-80v-96h80Zm64-128h104v32H328zm0 64h104v32H328zm0 64h104v32H328z"/>`),
		g.Group(children),
	)
}