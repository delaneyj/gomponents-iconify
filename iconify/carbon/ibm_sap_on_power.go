package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func IbmSapOnPower(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="M9 21H3c-1.103 0-2-.897-2-2v-6c0-1.103.897-2 2-2h6c1.103 0 2 .897 2 2v6c0 1.103-.897 2-2 2zm-6-8v6h6v-6H3zm8.693 16.325A14.05 14.05 0 0 1 4.2 23.538l1.685-1.078a12.043 12.043 0 0 0 6.421 4.96l-.614 1.904z"/><circle cx="19" cy="27" r="1" fill="currentColor"/><path fill="currentColor" d="M29 31H16c-1.103 0-2-.897-2-2v-4c0-1.103.897-2 2-2h13c1.103 0 2 .897 2 2v4c0 1.103-.897 2-2 2zm-13-6v4h13v-4H16zm10.89-4h2.184A14.07 14.07 0 0 0 30 16c0-.672-.064-1.338-.16-2h-2.026c.11.66.186 1.326.186 2c0 1.755-.387 3.43-1.11 5zm-1.39-9H19a1 1 0 0 1-1-1V4a1 1 0 0 1 1-1h12a1 1 0 0 1 .786 1.618l-5.5 7A1 1 0 0 1 25.5 12zM20 10h5.014l3.929-5H20v5zM5.886 9.54L4.201 8.46A13.955 13.955 0 0 1 16 2v2A11.96 11.96 0 0 0 5.886 9.54z"/>`),
		g.Group(children),
	)
}