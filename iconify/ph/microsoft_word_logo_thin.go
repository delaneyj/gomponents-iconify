package ph

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftWordLogoThin(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 256 256"),
		g.Raw(`<path fill="currentColor" d="M200 28H72a12 12 0 0 0-12 12v28H40a12 12 0 0 0-12 12v96a12 12 0 0 0 12 12h20v28a12 12 0 0 0 12 12h128a12 12 0 0 0 12-12V40a12 12 0 0 0-12-12Zm-44 72h48v56h-48ZM68 40a4 4 0 0 1 4-4h128a4 4 0 0 1 4 4v52h-48V80a12 12 0 0 0-12-12H68ZM36 176V80a4 4 0 0 1 4-4h104a4 4 0 0 1 4 4v96a4 4 0 0 1-4 4H40a4 4 0 0 1-4-4Zm164 44H72a4 4 0 0 1-4-4v-28h76a12 12 0 0 0 12-12v-12h48v52a4 4 0 0 1-4 4ZM72.12 153l-12-48a4 4 0 1 1 7.76-2l9.38 37.51l11.16-22.33a4 4 0 0 1 7.16 0l11.16 22.33l9.38-37.51a4 4 0 0 1 7.76 1.94l-12 48a4 4 0 0 1-3.44 3H108a4 4 0 0 1-3.58-2.21L92 128.94l-12.42 24.85a4 4 0 0 1-7.46-.82Z"/>`),
		g.Group(children),
	)
}