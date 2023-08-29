package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TonuinoNfcTools(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="18.422" height="35.002" x="16.647" y="8.498" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.647 39.502h-1.874a1.842 1.842 0 0 1-1.843-1.842V6.342A1.842 1.842 0 0 1 14.773 4.5H29.51a1.842 1.842 0 0 1 1.842 1.842v2.156m-14.705 3.943H35.07m-13.783-1.886h6.044"/><circle cx="30.12" cy="10.555" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.362 30.957l-8.186-4.726v9.452l8.186-4.726zm-2.112-6.83a3.729 3.729 0 0 0-4.764-.013v.013m7.038-2.585a7.332 7.332 0 0 0-9.331 0"/>`),
		g.Group(children),
	)
}