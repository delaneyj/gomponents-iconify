package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftOutlookLogoThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M88 100a28 28 0 1 0 28 28a28 28 0 0 0-28-28Zm0 48a20 20 0 1 1 20-20a20 20 0 0 1-20 20Zm128-40h-12V48a12 12 0 0 0-12-12h-80a12 12 0 0 0-12 12v20H40a12 12 0 0 0-12 12v96a12 12 0 0 0 12 12h36v20a12 12 0 0 0 12 12h128a12 12 0 0 0 12-12v-88a12 12 0 0 0-12-12Zm-57.17 56L220 119.84a.78.78 0 0 1 0 .16v88a1 1 0 0 1 0 .17Zm52.8-48l-7.63 5.51V116ZM108 48a4 4 0 0 1 4-4h80a4 4 0 0 1 4 4v79.29l-44 31.78l-4-2.89V80a12 12 0 0 0-12-12h-28ZM36 176V80a4 4 0 0 1 4-4h96a4 4 0 0 1 4 4v96a4 4 0 0 1-4 4H40a4 4 0 0 1-4-4Zm48 32v-20h52a12 12 0 0 0 12-12v-9.95L211.63 212H88a4 4 0 0 1-4-4Z"/>`),
		g.Group(children),
	)
}