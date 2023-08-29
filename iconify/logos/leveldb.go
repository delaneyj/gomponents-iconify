package logos

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Leveldb(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="#D4EB95" d="M128 72.17L0 135.367v128.535l128 64.268l128-64.268V135.367L128 72.17Z"/><path fill="#96DC75" d="M128 51.745L0 114.94v128.536l128 64.268l128-64.268V114.94L128 51.745Z"/><path fill="#317342" d="M128 0L0 63.197v128.535L128 256l128-64.268V63.197L128 0Z" opacity=".553"/><path fill="#34954C" d="M128 122.795L0 62.638v128.908L128 256l128-64.454V62.638l-128 60.157Z" opacity=".553"/>`),
		g.Group(children),
	)
}