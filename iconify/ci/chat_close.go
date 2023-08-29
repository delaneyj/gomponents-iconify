package ci

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ChatClose(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m10 13l2-2m0 0l2-2m-2 2l-2-2m2 2l2 2m-6.876 5.701L5.6 19.921c-.833.665-1.249.998-1.599.998a1 1 0 0 1-.783-.376C3 20.27 3 19.737 3 18.671V7.201c0-1.12 0-1.681.218-2.11c.192-.376.497-.681.874-.873C4.52 4 5.08 4 6.2 4h11.6c1.12 0 1.68 0 2.107.218c.377.192.683.497.875.874c.218.427.218.987.218 2.105v7.607c0 1.117 0 1.676-.218 2.104a2 2 0 0 1-.874.874c-.427.218-.987.218-2.105.218h-8.68c-.417 0-.624 0-.823.04a2.004 2.004 0 0 0-.508.179c-.18.091-.34.22-.657.474l-.011.008Z"/>`),
		g.Group(children),
	)
}