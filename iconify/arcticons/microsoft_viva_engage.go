package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func MicrosoftVivaEngage(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.002 30.384c0 7.794 5.04 12.116 11.01 12.116a8.832 8.832 0 1 0 0-17.664c-5.715 0-11.01-4.32-11.01-9.954v15.502Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.936 24.82a9.149 9.149 0 1 0-16.115-.048"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M6.002 14.882c0-5.57 3.537-9.382 8.832-9.382s7.045 3.348 8.552 5.825m.613 16.943a8.817 8.817 0 0 1 6.99-3.432c5.714 0 11.009-4.32 11.009-9.954v15.502c0 7.794-5.04 12.116-11.01 12.116A8.817 8.817 0 0 1 24 39.068m17.998-24.186c0-5.57-3.537-9.382-8.832-9.382c-5.294 0-7.045 3.348-8.552 5.825"/>`),
		g.Group(children),
	)
}