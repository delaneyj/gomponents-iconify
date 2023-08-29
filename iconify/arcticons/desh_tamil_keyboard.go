package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeshTamilKeyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M19.862 13.88H25.9v2.054a3.021 3.021 0 0 1-3.02 3.02h0a3.021 3.021 0 0 1-3.02-3.02v-2.055h0Zm9.124-3.891h-7.53v3.89"/><path d="M25.901 13.88h1.155a2.983 2.983 0 0 1 2.983 2.982v1.843a2.983 2.983 0 0 1-2.983 2.983H20.74c-1.784 0-2.384.6-2.778 1.301m7.939-9.11v-3.89"/></g><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.168 18.573H7.64c-1.148 0-2.078.93-2.078 2.078v19.771c0 1.148.93 2.078 2.078 2.078h32.72c1.148 0 2.078-.93 2.078-2.078V20.651c0-1.148-.93-2.078-2.078-2.078h-5.284"/><rect width="20.044" height="3.327" x="14.1" y="35.763" rx=".421" ry=".421"/><rect width="3.366" height="3.366" x="10.733" y="29.78" rx=".421" ry=".421"/><rect width="3.366" height="3.366" x="16.586" y="29.78" rx=".421" ry=".421"/><path d="M25.805 30.355v2.37a.42.42 0 0 1-.42.421H22.86a.42.42 0 0 1-.42-.42h0v-2.37"/><rect width="3.366" height="3.366" x="28.291" y="29.78" rx=".421" ry=".421"/><rect width="3.366" height="3.366" x="34.144" y="29.78" rx=".421" ry=".421"/><rect width="3.366" height="3.366" x="10.733" y="23.797" rx=".421" ry=".421"/><path d="M18.276 27.163h-1.27a.42.42 0 0 1-.42-.421v-.875m15.072 0v.875a.42.42 0 0 1-.421.42h-1.269"/><rect width="3.366" height="3.366" x="34.144" y="23.797" rx=".421" ry=".421"/><path d="m34.252 23.878l-10.13 7.768l-10.13-7.769a2.104 2.104 0 0 1-.824-1.67V7.184c0-.93.753-1.683 1.683-1.683h18.542c.93 0 1.683.754 1.683 1.683h0v15.025c0 .654-.304 1.271-.824 1.67Z"/></g>`),
		g.Group(children),
	)
}