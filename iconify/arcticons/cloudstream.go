package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Cloudstream(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.143 21.25v8.036l5.713-4.018Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.165 27.828H41.95c.687-2.529-1.202-3.652-2.858-3.652H37.49s.074-6.6-7.775-6.6c-1.222-2.062-3.548-3.408-6.935-3.408c-6.38 0-7.497 6.34-7.497 6.34a6.444 6.444 0 0 0-7.942 8.893H5.296a2.172 2.172 0 0 0 0 4.344c.017 0 .033-.004.05-.005l35.794.092a1.226 1.226 0 0 0 1.222-1.238l-.003-1.344h.806a1.711 1.711 0 1 0 0-3.422Z"/>`),
		g.Group(children),
	)
}