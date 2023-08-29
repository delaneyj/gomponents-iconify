package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UpfcmDistributor(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><ellipse cx="24" cy="38.527" rx="16.948" ry="4.973"/><path d="M24 4.5s-2.507-.056-2.757 3.39c0 0-7.285 1.661-7.266 12c.02 10.845-6.267 8.25-6.904 18.393M24 16.39l-4.898 6.556h2.335v4.195H24m5.976 6.732c0 3.345-2.676 6.056-5.976 6.056s-5.976-2.711-5.976-6.056M24 4.5s2.507-.056 2.757 3.39c0 0 7.285 1.661 7.266 12c-.02 10.845 6.267 8.25 6.904 18.393M24 16.39l4.898 6.556h-2.335v4.195H24"/></g>`),
		g.Group(children),
	)
}