package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AdobeAcrobat(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="4" ry="4"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M36.944 27.322c-2.013-2.136-7.581-1.27-8.894-1.051c-1.918-1.865-4.798-4.719-5.253-5.542c1.208-1.97 2.749-4.07 2.81-5.786c0-1.961-.788-4.063-2.96-4.063a2.232 2.232 0 0 0-1.838 1.068c-1.006 2.68.5 5.446 1.988 8.755c-.876 2.451-3.187 5.454-4.833 8.317c-2.232.875-6.855 2.679-7.424 5.48a2.031 2.031 0 0 0 .648 1.926c.53.468 1.22.715 1.926.692c2.854 0 2.705-4.631 4.85-8.098c1.611-.56 7.494-2.303 10.085-2.775c3.02 2.626 4.956 4.473 6.54 4.771a2.302 2.302 0 0 0 2.355-3.72v.026Z"/>`),
		g.Group(children),
	)
}