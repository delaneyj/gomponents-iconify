package ion

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-miterlimit="10" stroke-width="32" d="M416.07 272a160 160 0 1 0-160 160a160 160 0 0 0 160-160ZM142.12 91.21A46.67 46.67 0 0 0 112 80l-2.79.08C83.66 81.62 64 104 64.07 131c0 13.21 4.66 19.37 10.88 27.23a4.55 4.55 0 0 0 3.24 1.77h.88a3.23 3.23 0 0 0 2.54-1.31L142.38 99a5.38 5.38 0 0 0 1.55-4a5.26 5.26 0 0 0-1.81-3.79Zm227.76 0A46.67 46.67 0 0 1 400 80l2.79.08C428.34 81.62 448 104 447.93 131c0 13.21-4.66 19.37-10.88 27.23a4.55 4.55 0 0 1-3.24 1.76h-.88a3.23 3.23 0 0 1-2.54-1.31L369.62 99a5.38 5.38 0 0 1-1.55-4a5.26 5.26 0 0 1 1.81-3.79Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32" d="M256.07 160v112h-80m240 160l-40-40m-280 40l40-40"/>`),
		g.Group(children),
	)
}