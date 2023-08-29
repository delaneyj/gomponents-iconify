package tdesign

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LinkUnlink(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m4.304 2.89l3.535 3.536L6.425 7.84L2.89 4.304L4.303 2.89Zm5.949-1.709v4.5h-2v-4.5h2ZM1.347 8.09h4.5v2h-4.5v-2Zm13.482-2.454L11.785 8.68l-1.414-1.415l3.044-3.044a4.5 4.5 0 0 1 6.364 6.364l-3.044 3.044l-1.415-1.414l3.044-3.044a2.5 2.5 0 1 0-3.535-3.535Zm-9.192 9.192a2.5 2.5 0 1 0 3.535 3.536l3.044-3.044l1.414 1.414l-3.044 3.044a4.5 4.5 0 0 1-6.364-6.364l3.044-3.044l1.415 1.414l-3.044 3.044Zm12.678-1.081h4.5v2h-4.5v-2Zm-2.406 4.406v4.5h-2v-4.5h2Zm1.664-1.993l3.536 3.536l-1.415 1.414l-3.535-3.536l1.414-1.414Z"/>`),
		g.Group(children),
	)
}