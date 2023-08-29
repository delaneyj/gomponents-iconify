package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func DeshJapaneseKeyboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M13.168 18.573H7.64c-1.148 0-2.078.93-2.078 2.078v19.771c0 1.148.93 2.078 2.078 2.078h32.72c1.148 0 2.078-.93 2.078-2.078V20.651c0-1.148-.93-2.078-2.078-2.078h-5.284"/><rect width="20.044" height="3.327" x="14.1" y="35.763" rx=".421" ry=".421"/><rect width="3.366" height="3.366" x="10.733" y="29.78" rx=".421" ry=".421"/><rect width="3.366" height="3.366" x="16.586" y="29.78" rx=".421" ry=".421"/><path d="M25.805 30.355v2.37a.42.42 0 0 1-.42.421H22.86a.42.42 0 0 1-.42-.42h0v-2.37"/><rect width="3.366" height="3.366" x="28.291" y="29.78" rx=".421" ry=".421"/><rect width="3.366" height="3.366" x="34.144" y="29.78" rx=".421" ry=".421"/><rect width="3.366" height="3.366" x="10.733" y="23.797" rx=".421" ry=".421"/><path d="M18.276 27.163h-1.27a.42.42 0 0 1-.42-.421v-.875m15.072 0v.875a.42.42 0 0 1-.421.42h-1.269"/><rect width="3.366" height="3.366" x="34.144" y="23.797" rx=".421" ry=".421"/><path d="m34.252 23.878l-10.13 7.768l-10.13-7.769a2.104 2.104 0 0 1-.824-1.67V7.184c0-.93.753-1.683 1.683-1.683h18.542c.93 0 1.683.754 1.683 1.683h0v15.025c0 .654-.304 1.271-.824 1.67Z"/></g><g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"><path d="M23.324 9.99c-.734 4.614-.798 8.707.2 12.172"/><path d="M18.816 12.394a67.29 67.29 0 0 0 10.244-.476m-1.528 1.979c-1.393 4.094-3.235 7.23-6.112 8.065c-3.733 1.085-3.114-2.168-2.204-3.557c1.892-2.887 4.686-3.201 6.312-3.306c4.58-.294 5.636 3.652 4.559 5.26c-1.13 1.688-2.577 2.344-4.308 2.63"/></g>`),
		g.Group(children),
	)
}