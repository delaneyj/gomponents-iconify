package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Polytopia(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m29.995 10.322l11.354 6.821l-6.576 18.534l-8.721.201l-2.807-6.534l6.75-19.022zm-3.943 25.556l8.164-23.008M5.512 7.661L4.5 14.955l6.241 16.192l1.939.095l4.428-1.571l-6.304-16.824l-5.292-5.186z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m5.512 7.661l.548 7.168l6.62 16.413"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m7.536 32.412l10.837-2.319l1.518-1.94l1.367 2.735l-1.409 1.819l-11.217 2.319l-1.096-2.614l1.602-2.024l1.225-.221m9.528-2.014l-3.119.622"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17.977 33.094l1.914 4.884l-1.181 1.687l-4.048.674l-2.427-6.058m4.072-.667l2.403 6.051m22.639-22.522l2.151-.385l-6.688 18.59l-2.039.329m-4.778-25.355l2.909-.081L43.5 16.758M18.373 30.093l1.476 2.614"/>`),
		g.Group(children),
	)
}