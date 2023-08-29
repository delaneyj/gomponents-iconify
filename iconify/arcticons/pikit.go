package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Pikit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.026 16.713c-7.716 12.27-3.88 17.813-.812 20.449c3.072 2.699 13.36 4.48 17.813 4.608c4.48.128 14.178.23 16.723-3.8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.026 16.713C16.742 4.443 23.386 5.535 27.11 7.17s9.816 10.088 11.815 14.087c2.045 3.99 6.335 12.723 3.79 16.722"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.189 13.532C5.282 20.966 7.209 25.983 9.126 28.8c2.263 3.4 11.724 7.78 15.995 9.07c4.29 1.282 13.633 3.9 17.177.669"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.406 9.17c-4.663 4.726-4.108 9.034-3.154 11.815c1.309 3.871 9.361 10.542 13.178 12.905c3.817 2.354 12.178 7.307 16.359 5.09"/>`),
		g.Group(children),
	)
}