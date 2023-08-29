package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FilterX(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M40 16v37.828L109.024 136h41.791L76.896 48H459.51L304 242.388v158.985L241.373 464H240v-96h-32v128h46.627L336 414.627V253.612l160-200V16H40z"/><path fill="currentColor" d="m166.403 248.225l60.461-60.462l-22.627-22.628l-60.462 60.462l-60.462-60.462l-22.626 22.628l60.461 60.462l-60.461 60.462l22.626 22.627l60.462-60.462l60.462 60.462l22.627-22.627l-60.461-60.462z"/>`),
		g.Group(children),
	)
}