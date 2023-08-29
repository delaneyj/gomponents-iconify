package fluent_emoji_high_contrast

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Candle(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 32 32"),
		g.Raw(`<path fill="currentColor" d="m16.408 2.421l1.85 3.15a3.048 3.048 0 0 1-1.78 4.55V11.5H20.5a2 2 0 0 1 1.937 1.5h.063v6.75a1.75 1.75 0 0 1-1.5 1.732v6.106A2.412 2.412 0 0 1 18.588 30h-6.176A2.412 2.412 0 0 1 10 27.588v-8.673A1.5 1.5 0 0 1 9 17.5V13h.063A2 2 0 0 1 11 11.5h3.866v-1.38c-1.245-.347-2.188-1.477-2.24-2.818A3.035 3.035 0 0 1 13 5.717l.01-.017l.07-.12l1.856-3.159a.854.854 0 0 1 1.472 0ZM19 15h-.5v1a1.5 1.5 0 0 1-3 0v-1H12v12.588c0 .228.184.412.412.412h6.176a.412.412 0 0 0 .412-.412V15Zm-1.664-7.627a1.663 1.663 0 1 0-3.326 0a1.663 1.663 0 0 0 3.326 0Z"/>`),
		g.Group(children),
	)
}