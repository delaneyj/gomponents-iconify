package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Myntra(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<ellipse cx="28.311" cy="24" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="13.871" ry="4.22" transform="rotate(-68.92 28.311 24)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.727 24.215q.227.647.48 1.303c2.756 7.148 6.752 12.263 8.927 11.425s1.703-7.312-1.052-14.46s-6.752-12.264-8.927-11.426"/><ellipse cx="10.856" cy="24" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="13.871" ry="4.22" transform="rotate(-68.92 10.856 24)"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.272 24.215q.227.647.48 1.303c2.06 5.343 4.812 9.55 6.98 10.984M24 23.49q-.18-.501-.374-1.008c-2.756-7.148-6.752-12.263-8.927-11.425"/>`),
		g.Group(children),
	)
}