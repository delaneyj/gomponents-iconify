package si_glyph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BatteryEmpty(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 17 0"),
		g.Raw(`<path fill="currentColor" fill-rule="evenodd" d="M12.306 1h-1.307V.016H8V1H6.77C5.834 1 5 1.721 5 2.646v11.679C5 15.249 5.834 16 6.77 16h5.536c.936 0 1.694-.751 1.694-1.675V2.646C14 1.721 13.241 1 12.306 1zM13 14c0 .516-.484 1-1 1H7c-.515 0-1.011-.484-1.011-1V3c0-.514.496-1.051 1.011-1.051h5c.516 0 1 .536 1 1.051v11z"/>`),
		g.Group(children),
	)
}