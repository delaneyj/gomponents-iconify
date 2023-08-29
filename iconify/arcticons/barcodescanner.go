package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Barcodescanner(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M16.184 26.159H7.5v12.94h12.94v-3.617m-9.152-5.535h5.411v5.411h-5.411v-5.411Zm9.155-11.717V7.5H7.503v12.94h10.68m-6.893-9.15h5.412v5.412H11.29V11.29Zm24.175 9.152H39.1V7.502H26.16v8.686m3.788-4.899h5.411v5.412h-5.411v-5.412Z"/><circle cx="26.854" cy="26.853" r="10.588" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.5 34.5l10 9.999"/>`),
		g.Group(children),
	)
}