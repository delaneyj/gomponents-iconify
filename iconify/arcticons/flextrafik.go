package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Flextrafik(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M43.463 37.487V10.513a1.998 1.998 0 0 0-1.998-1.998H6.498A1.998 1.998 0 0 0 4.5 10.513v26.974a1.998 1.998 0 0 0 1.998 1.998h34.967a1.998 1.998 0 0 0 1.998-1.998Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m14.751 36.993l19.732.061M4.62 36.765l5.369.259m29.362.115l4.143.003s.138-4.04-.79-5.48c-.838-.953-3.2-1.355-4.887-2.308c-4.927-5.989-8.25-6.977-8.25-6.977l-24.948.019"/><circle cx="12.072" cy="37.024" r="2.084" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.837" cy="37.015" r="2.084" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><rect width="8.385" height="4.491" x="7.277" y="24.868" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".7"/><rect width="7.731" height="4.491" x="18.543" y="24.87" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx=".7"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.162 24.88h1.374a8.431 8.431 0 0 1 2.232 1.637l1.972 2.153c.262.286-.312.7-.7.7h-4.737a.699.699 0 0 1-.7-.7v-3.09c0-.389.172-.7.56-.7Zm-3.65-10.669l-4.45 5.777m4.45 0l-4.45-5.777m-1.925 4.677a2.233 2.233 0 0 1-1.933 1.1a2.203 2.203 0 0 1-2.225-2.18V16.39a2.203 2.203 0 0 1 2.225-2.18h0a2.203 2.203 0 0 1 2.225 2.18v.709h-4.45m-7.937-2.888h3.115m-1.796 5.777v-7.195a1.542 1.542 0 0 1 1.558-1.526h.606a1.853 1.853 0 0 1 1.584.638v6.991a1.101 1.101 0 0 0 1.113 1.09h.333"/>`),
		g.Group(children),
	)
}