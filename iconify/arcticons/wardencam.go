package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Wardencam(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 43.5c1.693 0 15.274-7.782 15.274-16.966V6.504C35.267 6.503 24 4.5 24 4.5S12.723 6.503 8.726 6.503v20.03C8.726 35.719 22.307 43.5 24 43.5Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.291 24.308s6.934 8.178 8.711 4.978M17.58 17.819l6.311-.978l11.822 8.267l-7.733.622l-1.067 4.711c-7.506-5.094-9.987-7.008-9.333-12.622Zm0 0l10.4 7.911"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.635 27.402l5.151-.327l-.727 3.205l-5.152.327l.728-3.205zm-16.019-.948a2.153 2.153 0 1 1 2.001-3.792m14.65-4.276a1.422 1.422 0 0 1 1.158.816m-1.605-3.299a3.724 3.724 0 0 1 4.142 2.916m-4.569-5.069a5.815 5.815 0 0 1 6.774 4.77"/>`),
		g.Group(children),
	)
}