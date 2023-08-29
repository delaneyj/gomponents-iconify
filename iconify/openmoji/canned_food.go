package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CannedFood(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<path fill="#d0cfce" d="M19 49h34v11H19z"/><path fill="#fff" d="M19 21h34v28H19z"/><path fill="#d0cfce" d="M19 4h34v16.795H19z"/><ellipse cx="36" cy="36.651" fill="#ea5a47" rx="8" ry="6.349"/><ellipse cx="36" cy="30.623" fill="#5c9e31" rx="3" ry="1.836"/><path fill="#9b9b9a" d="M47 4h6.5v17H47zm0 45h6v11h-6z"/><path fill="#d0cfce" d="M47 21h6v28h-6z"/><g fill="none" stroke="#000" stroke-linejoin="round" stroke-width="2"><path stroke-linecap="round" d="M40.67 31.49a6.311 6.311 0 0 1 3.368 5.39c0 3.65-3.592 6.595-8.019 6.595S28 40.531 28 36.88a6.316 6.316 0 0 1 3.24-5.3"/><path d="M39.387 30.46a2.38 2.38 0 0 1-.77 1.237a4.614 4.614 0 0 1-5.485 0a2.63 2.63 0 0 1-.77-1.237a3.529 3.529 0 0 1-.128-1.237a3.955 3.955 0 0 1 1.54.177a3.103 3.103 0 0 1 .833.412a3.406 3.406 0 0 1 1.155-1.855l.096-.088l.097.088a3.233 3.233 0 0 1 1.154 1.855a4.436 4.436 0 0 1 .834-.412a3.98 3.98 0 0 1 1.54-.177a2.524 2.524 0 0 1-.096 1.237Z"/><path stroke-linecap="round" d="M23 5h26M19 21h34M19 49h34M19.895 5H19v55h34V5h-.895"/></g>`),
		g.Group(children),
	)
}