package cil

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func QrCode(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M48 48h128V16H16v160h32V48z"/><path fill="currentColor" d="M176 176V80H80v96h96Zm-64-64h32v32h-32Zm216-64h136v128h32V16H328v32z"/><path fill="currentColor" d="M432 176V80h-96v96h96Zm-64-64h32v32h-32ZM176 464H48V336H16v160h160v-32z"/><path fill="currentColor" d="M176 336H80v96h96v-96Zm-32 64h-32v-32h32Zm320 64H328v32h168V336h-32v128z"/><path fill="currentColor" d="M272 304h128v64h32v-96H272v32z"/><path fill="currentColor" d="M432 432v-32H240V272H80v32h128v128h224zM208 80h32v96h-32z"/><path fill="currentColor" d="M80 240h224V80h-32v128H80v32zm256-32h96v32h-96zm0 128h32v32h-32zm-64 0h32v32h-32z"/>`),
		g.Group(children),
	)
}