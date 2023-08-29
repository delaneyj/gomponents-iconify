package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BluetoothKeyboardMouse(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M11.407 30.984v3.2h3.2v-3.2Zm5.498 0v3.2h3.2v-3.2Zm5.499 0v3.2h3.2v-3.2Zm5.499 0v3.2h3.2v-3.2Zm5.498 0v3.2h3.2v-3.2Zm-21.994-5.495v3.2h3.2v-3.2Zm5.498 0v3.2h3.2v-3.2Zm5.499 0v3.2h3.2v-3.2Zm5.499 0v3.2h3.2v-3.2Zm5.498 0v3.2h3.2v-3.2Zm-16.496 12.59h0a1.6 1.6 0 0 0 1.6 1.6h10.998a1.6 1.6 0 0 0 1.6-1.6h0a1.6 1.6 0 0 0-1.6-1.6H18.505a1.6 1.6 0 0 0-1.6 1.6Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.768 41.883V23.284a1.617 1.617 0 0 0-1.617-1.617H9.849a1.617 1.617 0 0 0-1.617 1.617v18.599A1.617 1.617 0 0 0 9.849 43.5h28.302a1.617 1.617 0 0 0 1.617-1.617ZM11.673 10.74A15.412 15.412 0 0 1 24 4.5h0a15.412 15.412 0 0 1 12.327 6.227m-20.229 3.306A9.918 9.918 0 0 1 24 9.905a9.918 9.918 0 0 1 7.92 4.14m-11.713 3.098a4.754 4.754 0 0 1 7.586 0"/>`),
		g.Group(children),
	)
}