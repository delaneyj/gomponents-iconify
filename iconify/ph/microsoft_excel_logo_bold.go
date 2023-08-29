package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftExcelLogoBold(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 20H72a20 20 0 0 0-20 20v16H36a20 20 0 0 0-20 20v104a20 20 0 0 0 20 20h16v16a20 20 0 0 0 20 20h128a20 20 0 0 0 20-20V40a20 20 0 0 0-20-20Zm-32 88h28v40h-28Zm28-24h-28v-8a20 20 0 0 0-20-20V44h48ZM76 44h48v12H76ZM40 80h104v96H40Zm36 120h48v12H76Zm72 12v-12a20 20 0 0 0 20-20v-8h28v40Zm-83.68-50.78a12 12 0 0 1-1.54-16.9L76.38 128l-13.6-16.32a12 12 0 1 1 18.44-15.36L92 109.25l10.78-12.93a12 12 0 0 1 18.44 15.36L107.62 128l13.6 16.32a12 12 0 1 1-18.44 15.36L92 146.75l-10.78 12.93a12 12 0 0 1-16.9 1.54Z"/>`),
		g.Group(children),
	)
}