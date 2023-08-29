package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Medicinkortet(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M21.37 23.278c2.003 1.4.954 5.091-4.547 9.585c0 0-4.29-3.249-4.364-6.386s3.142-3.058 4.358.42c0 0 2.551-5.02 4.553-3.62Z"/><rect width="16.706" height="8.563" x="8.464" y="5.592" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="1.446"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m23.371 14.492l.043 1.925c.126 2.223 4.683 2.616 4.683 5.789v13.827M24.2 39.932H9.137a3.593 3.593 0 0 1-3.601-3.601V22.206c0-2.824 4.08-3.207 4.404-5.745l.064-1.608m26.627 10.601c1.609 1.25 3.232 2.48 4.831 3.742a3.232 3.232 0 0 1 .528 4.155a3.063 3.063 0 0 1-3.9 1.077a19.545 19.545 0 0 1-1.918-1.401c-1.15-.895-2.313-1.771-3.453-2.678a3.232 3.232 0 0 1-.528-4.155a3.063 3.063 0 0 1 3.9-1.076a3.153 3.153 0 0 1 .54.336Zm-9.26 10.718c2.037.008 4.074-.007 6.11.016a3.232 3.232 0 0 1 2.948 2.89a3.066 3.066 0 0 1-2.275 3.279a11.871 11.871 0 0 1-2.257.095c-1.546-.007-3.092.009-4.638-.014a3.232 3.232 0 0 1-2.947-2.89a3.066 3.066 0 0 1 2.275-3.278a3.142 3.142 0 0 1 .784-.098Zm2.929 6.245l.008-6.053m5.022-4.232l3.682-4.804"/>`),
		g.Group(children),
	)
}