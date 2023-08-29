package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Materialpods(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.7 17.056a7.345 7.345 0 0 1 3.24 5.98V40.82a1.68 1.68 0 0 0 1.678 1.679h2.439a1.68 1.68 0 0 0 1.679-1.68V19.69c0-4.4-.964-8.777-3.35-11.446a8.944 8.944 0 0 0-.727-.696C14.544 4.042 9.586 5.232 5.6 9.638a1.537 1.537 0 0 0-.336.885v6.593a1.624 1.624 0 0 0 .326.897c3.15 3.711 6.17 4.852 11.313 4.32"/><circle cx="9.447" cy="13.933" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.736 39.15h-3.26M15.628 9.964a8.58 8.58 0 0 1 2.338 3.342m16.334 3.75a7.345 7.345 0 0 0-3.24 5.98V40.82a1.68 1.68 0 0 1-1.678 1.679h-2.439a1.68 1.68 0 0 1-1.679-1.68V19.69c0-4.4.964-8.777 3.35-11.446a8.944 8.944 0 0 1 .727-.696c4.115-3.506 9.073-2.316 13.059 2.09a1.537 1.537 0 0 1 .336.885v6.593a1.624 1.624 0 0 1-.326.897c-3.15 3.711-6.17 4.852-11.313 4.32"/><circle cx="38.553" cy="13.933" r=".75" fill="currentColor"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.264 39.15h3.26m3.848-29.186a8.58 8.58 0 0 0-2.337 3.342"/>`),
		g.Group(children),
	)
}