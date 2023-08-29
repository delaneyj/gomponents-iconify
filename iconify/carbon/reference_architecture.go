package carbon

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func ReferenceArchitecture(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<circle cx="18" cy="26" r="1" fill="currentColor"/><path fill="currentColor" d="M28 20c1.103 0 2-.897 2-2v-4c0-1.103-.897-2-2-2h-1V7c0-1.103-.897-2-2-2h-5V4c0-1.103-.897-2-2-2h-4c-1.103 0-2 .897-2 2v1H7c-1.103 0-2 .897-2 2v5.142c-1.72.447-3 2-3 3.858s1.28 3.41 3 3.858V25c0 1.103.897 2 2 2h7v1c0 1.103.897 2 2 2h12c1.103 0 2-.897 2-2v-4c0-1.103-.897-2-2-2h-1v-2h1Zm0-2h-4v-4h4v4ZM14 4h4v4h-4V4Zm-2 3v1c0 1.103.897 2 2 2h4c1.103 0 2-.897 2-2V7h5v5h-1c-1.103 0-2 .897-2 2v1H9.858A3.994 3.994 0 0 0 7 12.142V7h5Zm-8 9c0-1.103.897-2 2-2s2 .897 2 2s-.897 2-2 2s-2-.897-2-2Zm24 12H16v-4h12v4Zm-3-6h-9c-1.103 0-2 .897-2 2v1H7v-5.142A3.994 3.994 0 0 0 9.858 17H22v1c0 1.103.897 2 2 2h1v2Z"/>`),
		g.Group(children),
	)
}