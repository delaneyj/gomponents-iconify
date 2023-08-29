package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Florisboard(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.168 18.573H7.64a2.078 2.078 0 0 0-2.078 2.078v19.771A2.078 2.078 0 0 0 7.64 42.5h32.72a2.078 2.078 0 0 0 2.078-2.078V20.651a2.078 2.078 0 0 0-2.078-2.078h-5.284"/><rect width="20.044" height="3.327" x="14.1" y="35.763" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".421"/><rect width="3.366" height="3.366" x="10.733" y="29.78" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".421"/><rect width="3.366" height="3.366" x="16.586" y="29.78" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".421"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.805 30.355v2.37a.42.42 0 0 1-.42.421H22.86a.42.42 0 0 1-.42-.42v-2.37"/><rect width="3.366" height="3.366" x="28.291" y="29.78" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".421"/><rect width="3.366" height="3.366" x="34.144" y="29.78" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".421"/><rect width="3.366" height="3.366" x="10.733" y="23.797" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".421"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M18.276 27.163h-1.27a.42.42 0 0 1-.42-.421v-.875m15.072 0v.875a.42.42 0 0 1-.421.42h-1.269"/><rect width="3.366" height="3.366" x="34.144" y="23.797" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".421"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m34.252 23.878l-10.13 7.768l-10.13-7.768a2.104 2.104 0 0 1-.824-1.67V7.183A1.683 1.683 0 0 1 14.85 5.5h18.542a1.683 1.683 0 0 1 1.683 1.683v15.025a2.104 2.104 0 0 1-.824 1.67ZM21.75 9.989h6.5m-6.5 6.5h4.225m-4.225-6.5v13"/>`),
		g.Group(children),
	)
}