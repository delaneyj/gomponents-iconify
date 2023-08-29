package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Frameo(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M42.5 16.887a20.701 20.701 0 0 0-2.295 2.641a7.235 7.235 0 0 0-1.225 2.852c-.16 1.283.124 3.883-2.524 4.515a3.179 3.179 0 0 1-3.66-1.95a10.18 10.18 0 0 1-.23-4.014a11.376 11.376 0 0 1 .612-3.798s-11.382 1.231-15.183 1.57c-1.012.09-3.406.253-3.612-2.46s1.941-3.07 2.952-3.2c3.375-.435 13.26-1.58 13.26-1.58"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.364 8.54s-4.593.59-6.275.772c-1.35.146-2.164.623-2.248 1.548a1.932 1.932 0 0 0 .637 1.555"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M26.791 5.631C24 7.047 25.628 9.133 25.628 9.133M5.5 31.244a20.701 20.701 0 0 0 2.295-2.64a7.234 7.234 0 0 0 1.225-2.852c.16-1.284-.124-3.883 2.524-4.516a3.179 3.179 0 0 1 3.66 1.95a10.18 10.18 0 0 1 .23 4.015a11.375 11.375 0 0 1-.612 3.798s11.382-1.231 15.183-1.57c1.012-.09 3.406-.253 3.612 2.46s-1.941 3.07-2.952 3.2c-3.375.435-13.26 1.58-13.26 1.58"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M17.636 39.592s4.593-.59 6.275-.772c1.35-.146 2.164-.623 2.248-1.548a1.932 1.932 0 0 0-.637-1.556"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.209 42.5c2.791-1.415 1.164-3.501 1.164-3.501"/><rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2"/>`),
		g.Group(children),
	)
}