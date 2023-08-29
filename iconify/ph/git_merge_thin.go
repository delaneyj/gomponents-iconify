package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GitMergeThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M208 116a28 28 0 0 0-27.71 24H152a44.2 44.2 0 0 1-35.2-17.6L87.28 83a28 28 0 1 0-11.28.71v88.58a28 28 0 1 0 8 0V92l26.4 35.2A52.26 52.26 0 0 0 152 148h28.29A28 28 0 1 0 208 116ZM60 56a20 20 0 1 1 20 20a20 20 0 0 1-20-20Zm40 144a20 20 0 1 1-20-20a20 20 0 0 1 20 20Zm108-36a20 20 0 1 1 20-20a20 20 0 0 1-20 20Z"/>`),
		g.Group(children),
	)
}