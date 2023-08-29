package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Reboot(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M27.55 34.318c1.79-2.404 6.18-7.911 5.816-13.324c0 0-.06-.886-.071-1.482a17.988 17.988 0 0 1 4.104-11.707a13.078 13.078 0 0 0-7.725 9.103l-.05.238a21.253 21.253 0 0 1-13.29-12.931q-.126-.355-.238-.715a18.043 18.043 0 0 0 1.483 6.435a22.51 22.51 0 0 0 5.103 7.486c1.301 1.175 3.288 2.887 4.527 4.121"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.828 25.544q-.116-.194-.227-.39a10.116 10.116 0 0 0 12.213 1.375a12.968 12.968 0 0 0 4.395-4.986c1.264 1.47 1.639 3.94 1.898 5.863a10.174 10.174 0 0 1-.879 5.67c-.132.28-.277.554-.43.823c-1.445 2.533-3.601 4.605-5.04 7.142a14.947 14.947 0 0 0-1.424 3.46a17.662 17.662 0 0 1 3.955-13.638q.117-.139.237-.275a10.346 10.346 0 0 1-8.172.589a12.67 12.67 0 0 1-6.333-5.32q-.098-.155-.192-.312Z"/><ellipse cx="27.209" cy="11.277" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.526" ry="2.901" transform="rotate(-23.234 27.209 11.277)"/>`),
		g.Group(children),
	)
}