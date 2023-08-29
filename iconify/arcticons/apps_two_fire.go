package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AppsTwoFire(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.288 22.3c1.11-.45 3.09-1.05 3.69-.33s-.17 2.48-.92 3.8"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.808 22.82c1.76 1.4 7 3.54 12.49 3.54a17 17 0 0 0 10.16-3.08"/><circle cx="21.931" cy="30.878" r=".7" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.931 33.328v5.299m-4.131 0v-6.6a1.4 1.4 0 0 1 1.4-1.4h0a1.63 1.63 0 0 1 1.42.59m-3.999 2.111h2.8m4.655 2a2 2 0 0 1 2-2h0m-2 0v5.299m7.044-1a2 2 0 0 1-1.74 1h0a2 2 0 0 1-2-2v-1.29a2 2 0 0 1 2-2h0a2 2 0 0 1 2 2v.65h-4m-9.984-20.783a6.604 6.604 0 0 1 13.207 0v3.116a.691.691 0 0 1-.678.688H18.093a.692.692 0 0 1-.696-.688h0Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" d="M28.727 10.593L30.5 8.821m-11.227 1.772L17.5 8.821"/><circle cx="21.144" cy="14.158" r=".75" fill="currentColor"/><circle cx="26.867" cy="14.158" r=".75" fill="currentColor"/>`),
		g.Group(children),
	)
}