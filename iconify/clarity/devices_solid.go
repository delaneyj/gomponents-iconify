package clarity

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DevicesSolid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 36 36"),
		g.Raw(`<path fill="currentColor" d="M32 13h-8a2 2 0 0 0-2 2v15a2 2 0 0 0 2 2h8a2 2 0 0 0 2-2V15a2 2 0 0 0-2-2Zm0 2v13h-8V15Z" class="clr-i-solid clr-i-solid-path-1"/><path fill="currentColor" d="M28 4H4a2 2 0 0 0-2 2v16a2 2 0 0 0 2 2h8v2H9.32A1.2 1.2 0 0 0 8 27a1.2 1.2 0 0 0 1.32 1h10.6v-.37H20V22H4V6h24v5h2V6a2 2 0 0 0-2-2Z" class="clr-i-solid clr-i-solid-path-2"/><path fill="none" d="M0 0h36v36H0z"/>`),
		g.Group(children),
	)
}