package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AndromedaKwgt(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.494 27.656L24 18.219L12.437 42.5H5.499L24 5.5l10.94 21.877m4.364 8.729L42.5 42.5h-6.938l-2.839-5.962"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.112 29.679l-1.15 4.599l-1.149-4.599l-1.15 4.599l-1.15-4.599m8.952 1.554a1.557 1.557 0 0 0-1.609-1.55a1.618 1.618 0 0 0-1.436 1.608v1.436a1.545 1.545 0 0 0 1.551 1.551h0a1.545 1.545 0 0 0 1.551-1.551h-1.55m2.332-3.048h3.047m-1.496 4.599v-4.599m-15.121 0v4.599m2.47-.001l-1.897-2.3l1.897-2.299m-1.897 2.301h-.573"/>`),
		g.Group(children),
	)
}