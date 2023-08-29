package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func BestConnect(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.567 13.869A14.46 14.46 0 0 1 43.5 19.125h0a14.461 14.461 0 0 1-11.615 14.177"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.758 42.83a7.91 7.91 0 1 1 10.371-7.513h0a7.906 7.906 0 0 1-5.522 7.535"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.758 39.494a4.921 4.921 0 1 1 7.4-4.252h0a4.926 4.926 0 0 1-2.551 4.317m-2.388 3.778v-6.648M12.41 34.434a10.403 10.403 0 0 1 2.49-20.503h.007a10.404 10.404 0 0 1 10.408 10.4v.003h0a10.403 10.403 0 0 1-.474 3.109"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M12.41 30.125a6.318 6.318 0 1 1 8.787-5.823h0a6.322 6.322 0 0 1-3.934 5.851m-2.384 4.604V26.11m5.916-10.37a8.91 8.91 0 0 1 17.153 3.384h0a8.91 8.91 0 0 1-6.792 8.652m-2.124-6.224v7.085"/><circle cx="24.219" cy="35.22" r="1.296" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="14.879" cy="24.32" r="1.686" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="29.035" cy="19.084" r="2.475" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}