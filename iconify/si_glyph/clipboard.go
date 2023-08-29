package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Clipboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M11.002 1.016v.937h3.014v13.089H2.971V1.953h2.982v-.906H2V16h13V1.016h-3.998z"/><path d="M9.95 1.5C9.95.672 9.298 0 8.491 0c-.808 0-1.46.672-1.46 1.5c0 .176.035.343.09.5h-.057v1h2.873V2h-.076c.054-.157.089-.324.089-.5zm-1.981.516V.959h1.047v1.057H7.969zM11.031 3v1.016H5.969V3h-1.94v11h8.951V3h-1.949z"/></g>`),
		g.Group(children),
	)
}