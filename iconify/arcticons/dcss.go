package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Dcss(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M38.899 15.831a2.273 2.273 0 0 1 .694 1.54c0 2.87-6.981 5.197-15.593 5.197S8.407 20.241 8.407 17.37S15.388 12.172 24 12.172a42.032 42.032 0 0 1 7.585.655"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.475 17.856a5.852 5.852 0 0 0 .706 3.825a1.11 1.11 0 0 1 .172.644m-.282 12.25c2.216 2.943 5.96 4.39 14.929 4.39"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M10.911 36.386A3.303 3.303 0 0 0 9.572 38.6c.021 1.21-1.884 1.49-1.568 2.355a2.159 2.159 0 0 0 2.408 1.546c1.586 0 1.764-1.379 1.804-2.084s1.628-.826 1.72-2.563m9.918-15.287s-.325 2.902.938 3.052c1.28.152.96 3.227 1.437 4.312c.949 2.157 3.078-1.157 3.074-3.898c-.004-2.695.696-1.168.944-3.9m6.841 14.253a3.302 3.302 0 0 1 1.34 2.213c-.022 1.21 1.884 1.49 1.567 2.355a2.158 2.158 0 0 1-2.407 1.546c-1.587 0-1.765-1.379-1.805-2.084s-1.627-.826-1.72-2.563M9.845 19.551c1.168-2.243 7.06-3.95 14.155-3.95a38.753 38.753 0 0 1 7.09.62m6.245 2.371a3.23 3.23 0 0 1 .804.93m-5.465-1.662s4.158-12.238 7.624-12.357s1.673 4.143.231 5.303s-5.388 8.22-5.388 8.22a1.72 1.72 0 0 1-2.467-1.166Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M39.525 17.856a5.852 5.852 0 0 1-.707 3.825a1.11 1.11 0 0 0-.171.644m.282 12.25c-2.216 2.944-5.961 4.39-14.929 4.39"/><ellipse cx="39.33" cy="28.259" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2.979" ry="3.913" transform="rotate(-23.695 39.33 28.26)"/><ellipse cx="8.67" cy="28.26" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="3.913" ry="2.979" transform="rotate(-66.305 8.67 28.26)"/>`),
		g.Group(children),
	)
}