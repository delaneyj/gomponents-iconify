package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Planesandroid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="7.852" height="7.852" x="25.602" y="9.251" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".274" transform="rotate(-55.994 29.528 13.177)"/><rect width="7.852" height="7.852" x="21.211" y="15.76" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".274" transform="rotate(-55.994 25.137 19.686)"/><rect width="7.852" height="7.852" x="27.72" y="20.152" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".274" transform="rotate(-55.994 31.646 24.078)"/><rect width="7.852" height="7.852" x="16.819" y="22.27" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".274" transform="rotate(-55.994 20.745 26.196)"/><rect width="7.852" height="7.852" x="12.428" y="28.779" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".274" transform="rotate(-55.994 16.354 32.705)"/><rect width="7.852" height="7.852" x="18.937" y="33.171" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".274" transform="rotate(-55.994 22.863 37.097)"/><rect width="7.852" height="7.852" x="5.918" y="24.387" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".274" transform="rotate(-55.994 9.844 28.314)"/><rect width="7.852" height="7.852" x="14.701" y="11.369" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".274" transform="rotate(-55.994 18.628 15.295)"/><rect width="7.852" height="7.852" x="34.23" y="24.544" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".274" transform="rotate(-55.994 38.156 28.47)"/><rect width="7.852" height="7.852" x="8.192" y="6.977" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".274" transform="rotate(-55.994 12.118 10.903)"/>`),
		g.Group(children),
	)
}