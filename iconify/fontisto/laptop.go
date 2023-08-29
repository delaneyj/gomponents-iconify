package fontisto

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Laptop(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<path fill="currentColor" d="M2.998 18.833H24.38c.843 0 1.526-.683 1.526-1.526V2.034a2.035 2.035 0 0 0-2.035-2.035H3.507a2.035 2.035 0 0 0-2.035 2.035v15.274a1.529 1.529 0 0 0 1.526 1.527zM4.02 2.549h19.336v13.736H4.02zm22.845 17.22H.509a.51.51 0 0 0-.509.509v1.886a2.097 2.097 0 0 0 2.295 1.825l-.009.001h22.801a2.094 2.094 0 0 0 2.285-1.815l.001-.01v-1.886a.51.51 0 0 0-.509-.509zm-9.996 2.62h-6.383a.51.51 0 1 1 0-1.02h.021h-.001h6.38a.51.51 0 1 1 0 1.02h-.021h.001z"/>`),
		g.Group(children),
	)
}