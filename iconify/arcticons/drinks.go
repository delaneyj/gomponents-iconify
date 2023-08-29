package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Drinks(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.73 7.29a68.61 68.61 0 0 0-1.18 8.2c0 3.25 2.12 6 4.9 7.67a11.62 11.62 0 0 0 4.46 1.53v14.15c-3.74.16-6.62 1.09-6.62 2.32S19.75 43.5 24 43.5s7.73-1.08 7.73-2.34a.9.9 0 0 0-.13-.44c-.6-1-3.22-1.74-6.5-1.88V24.69c4.61-.54 9.33-4.29 9.33-9.2a74.84 74.84 0 0 0-1.14-8.18m1 9.51a10.78 10.78 0 0 1-1.39.32a13.28 13.28 0 0 1-8.06-.92a23.09 23.09 0 0 0-11.2-1.73"/><ellipse cx="24.01" cy="7.29" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="9.28" ry="2.57"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.91 24.69s.81.08 1.08.08s1.11-.08 1.11-.08"/>`),
		g.Group(children),
	)
}