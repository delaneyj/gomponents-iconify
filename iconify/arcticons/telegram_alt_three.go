package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func TelegramAltThree(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.714 28.155l7.62-6.553c1.282-1.101.6-3.202-1.085-3.34l-10.016-.821a1.903 1.903 0 0 1-1.6-1.163l-3.877-9.273c-.652-1.558-2.86-1.558-3.512 0l-3.877 9.273a1.903 1.903 0 0 1-1.6 1.163l-10.017.821c-1.684.138-2.366 2.239-1.085 3.34l4.628 3.98a4.11 4.11 0 0 0 3.127.969l10.362-1.135c.852-.093 1.205 1.06.447 1.46l-9.301 4.91a4.11 4.11 0 0 0-2.081 2.69l-1.264 5.34c-.389 1.644 1.398 2.943 2.841 2.065l8.587-5.223c.607-.37 1.37-.37 1.978 0l8.586 5.223c1.444.878 3.23-.42 2.842-2.065l-2.314-9.78a1.903 1.903 0 0 1 .611-1.881Z"/>`),
		g.Group(children),
	)
}