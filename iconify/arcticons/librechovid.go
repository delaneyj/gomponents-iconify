package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Librechovid(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.96 17.59v10.15m5.16-5.08H18.81M24 35.46c10.79-5.31 10.4-17 10.4-17v-3.93A32.33 32.33 0 0 0 24 12.91a32.33 32.33 0 0 0-10.4 1.62v3.92s-.4 11.7 10.4 17.01Zm-.01-22.55V7.89"/><ellipse cx="23.99" cy="5.63" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.29" ry="2.25"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m32.77 14.03l1.83-1.79"/><ellipse cx="36.22" cy="10.64" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.29" ry="2.25"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M33.87 22.69h5.11"/><ellipse cx="41.27" cy="22.69" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.29" ry="2.25"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m30.92 29.55l3.64 3.58"/><ellipse cx="36.18" cy="34.73" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.29" ry="2.25"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.93 35.46v2.19"/><ellipse cx="23.93" cy="39.69" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.08" ry="2.04"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m17 29.47l-3.82 3.77"/><ellipse cx="11.71" cy="34.69" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.08" ry="2.05"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.05 22.63H8.76"/><ellipse cx="6.66" cy="22.63" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.1" ry="2.07"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m15.23 14.03l-1.96-1.93"/><ellipse cx="11.75" cy="10.6" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.15" ry="2.11"/>`),
		g.Group(children),
	)
}