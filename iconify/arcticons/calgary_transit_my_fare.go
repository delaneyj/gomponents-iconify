package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CalgaryTransitMyFare(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.478 8.595L34.25 6.51a2.536 2.536 0 0 1 3.124 1.831l5.046 19.905a2.623 2.623 0 0 1-1.857 3.169L32.792 33.5a2.536 2.536 0 0 1-3.124-1.832l-5.046-19.905a2.622 2.622 0 0 1 1.856-3.168ZM37.32 29.317l-2.423.65"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.127 39.157v-20.44c0-5 8.155-4.22 11.428-4.22h5.76"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.968 17.076H12.991c-2.493 0-2.822.326-2.822 2.823v10.548s-.371 1.63 9.976 1.63c5.694 0 8.11-.466 9.13-.884m-19.014-10.72H26.69"/><rect width="3.741" height="2.423" x="9.74" y="39.157" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".429" ry=".429"/><rect width="3.741" height="2.423" x="9.74" y="39.157" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".429" ry=".429"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.246 35.181h-3.709v-1.534c0-.66 3.566-.262 3.71 1.534Zm11.744 0h3.709v-1.534c0-.66-3.566-.262-3.71 1.534Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.127 39.157h23.016v-5.765M7.768 25.435H5.5v3.451c0 2.197 2.58 1.948 2.58 1.948"/><rect width="3.741" height="2.423" x="26.282" y="39.157" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".429" ry=".429"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.111 22.238c.705.497 1.43.646 2.39.389l1.28-.343a2.215 2.215 0 0 0 1.56-2.704h0a2.214 2.214 0 0 0-2.703-1.562l-1.387.372a2.214 2.214 0 0 1-2.703-1.561h0a2.214 2.214 0 0 1 1.56-2.705l1.28-.342c.96-.258 1.657-.216 2.39.388m-3.612-2.222l.512 2.023m3.045 10.489l-.513-2.023"/>`),
		g.Group(children),
	)
}