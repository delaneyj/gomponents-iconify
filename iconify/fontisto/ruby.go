package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Ruby(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="m6.28 7.906l-.014-.014l-2.96 2.96l7.186 7.174l2.96-2.946l4.226-4.226l-2.96-2.96V7.88H6.266z"/><path fill="currentColor" d="M10.466 0L0 6v12l10.466 6l10.466-6V6zm8.466 16.854l-8.466 4.88L2 16.854V7.12l8.466-4.88l8.466 4.88z"/>`),
		g.Group(children),
	)
}