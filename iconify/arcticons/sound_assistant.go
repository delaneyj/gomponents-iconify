package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func SoundAssistant(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.321 10.786a5.286 5.286 0 1 1-10.571 0a5.286 5.286 0 0 1 10.571 0ZM37.214 24a5.286 5.286 0 1 1-10.571 0a5.286 5.286 0 0 1 10.571 0Zm-7.928 13.214a5.286 5.286 0 1 1-10.572 0a5.286 5.286 0 0 1 10.572 0Zm0 0h12.553m-35.678 0h12.553M37.214 24h4.625M6.161 24h20.482m-1.322-13.214h16.518m-35.678 0h8.589"/>`),
		g.Group(children),
	)
}