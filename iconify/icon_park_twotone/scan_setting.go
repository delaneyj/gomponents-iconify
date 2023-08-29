package icon_park_twotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ScanSetting(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<mask id="ipTScanSetting0"><g fill="none" stroke="#fff" stroke-linecap="round" stroke-linejoin="round" stroke-width="4"><rect width="38" height="38" x="5" y="5" fill="#555" rx="3"/><path d="m34 21l3 3l-3 3m-20-6l-3 3l3 3m13-13l-3-3l-3 3m6 20l-3 3l-3-3"/></g></mask><path fill="currentColor" d="M0 0h48v48H0z" mask="url(#ipTScanSetting0)"/>`),
		g.Group(children),
	)
}