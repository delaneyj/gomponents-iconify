package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ClipboardChecked(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<g fill="currentColor" fill-rule="evenodd"><path d="M12.002 1.016v.937h3.014v13.089H3.971V1.953h2.982v-.906H3V16h13V1.016h-3.998z"/><path d="M10.95 1.5c0-.828-.652-1.5-1.459-1.5c-.808 0-1.46.672-1.46 1.5c0 .176.035.343.09.5h-.057v1h2.873V2h-.076c.054-.157.089-.324.089-.5zm-1.981.516V.959h1.047v1.057H8.969zm3.062 2H6.969V3H5.03v11h8.951V3h-1.95v1.016zM7.453 8.127l.762-.762l1.414 1.414l2.088-2.09l.811.811l-2.851 2.852l-2.224-2.225z"/></g>`),
		g.Group(children),
	)
}