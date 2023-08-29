package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func FileJpgBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M120 140h-16a12 12 0 0 0-12 12v56a12 12 0 0 0 24 0v-4h4a32 32 0 0 0 0-64Zm0 40h-4v-16h4a8 8 0 0 1 0 16Zm104 20.87a12 12 0 0 1-3.33 8.3A34 34 0 0 1 196 220c-19.85 0-36-17.94-36-40s16.15-40 36-40a33.39 33.39 0 0 1 18.71 5.8a12 12 0 1 1-13.42 19.89A9.39 9.39 0 0 0 196 164c-6.5 0-12 7.33-12 16s5.5 16 12 16a9 9 0 0 0 2.7-.42A12 12 0 0 1 208 176h4a12 12 0 0 1 12 12ZM80 152v38a30 30 0 0 1-60 0a12 12 0 0 1 24 0a6 6 0 0 0 12 0v-38a12 12 0 0 1 24 0Zm-32-32a12 12 0 0 0 12-12V44h76v48a12 12 0 0 0 12 12h48v4a12 12 0 0 0 24 0V88a12 12 0 0 0-3.51-8.48l-56-56A12 12 0 0 0 152 20H56a20 20 0 0 0-20 20v68a12 12 0 0 0 12 12Zm112-63l23 23h-23Z"/>`),
		g.Group(children),
	)
}