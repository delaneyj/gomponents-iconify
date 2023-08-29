package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OneCKeyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.168 18.573H7.64c-1.148 0-2.078.93-2.078 2.078v19.771c0 1.148.93 2.078 2.078 2.078h32.72c1.148 0 2.078-.93 2.078-2.078V20.651c0-1.148-.93-2.078-2.078-2.078h-5.284"/><rect width="20.044" height="3.327" x="14.1" y="35.763" rx=".421" ry=".421"/><rect width="3.366" height="3.366" x="10.733" y="29.78" rx=".421" ry=".421"/><rect width="3.366" height="3.366" x="16.586" y="29.78" rx=".421" ry=".421"/><path d="M25.805 30.355v2.37a.42.42 0 0 1-.42.421H22.86a.42.42 0 0 1-.42-.42h0v-2.37"/><rect width="3.366" height="3.366" x="28.291" y="29.78" rx=".421" ry=".421"/><rect width="3.366" height="3.366" x="34.144" y="29.78" rx=".421" ry=".421"/><rect width="3.366" height="3.366" x="10.733" y="23.797" rx=".421" ry=".421"/><path d="M18.276 27.163h-1.27a.42.42 0 0 1-.42-.421v-.875m15.072 0v.875a.42.42 0 0 1-.421.42h-1.269"/><rect width="3.366" height="3.366" x="34.144" y="23.797" rx=".421" ry=".421"/><path d="m34.252 23.878l-10.13 7.768l-10.13-7.769a2.104 2.104 0 0 1-.824-1.67V7.184c0-.93.753-1.683 1.683-1.683h18.542c.93 0 1.683.754 1.683 1.683h0v15.025c0 .654-.304 1.271-.824 1.67Z"/></g><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m16.257 11.759l3.25-1.77v13m12.236-4.359v.053a4.306 4.306 0 0 1-4.306 4.306h0a4.306 4.306 0 0 1-4.307-4.306v-4.387a4.306 4.306 0 0 1 4.307-4.307h0a4.306 4.306 0 0 1 4.306 4.307v.053"/>`),
		g.Group(children),
	)
}