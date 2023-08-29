package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Osmtracker(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="31.953" cy="18.152" r="10.548" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M14.157 39.41a15.11 15.11 0 0 1 1.836.57a49.547 49.547 0 0 0 6.263-2.066a77.714 77.714 0 0 0 8.464 2.438a24.38 24.38 0 0 0 8.531-2.235a23.632 23.632 0 0 1 2.179-8.463a11.426 11.426 0 0 0-1.342-4.788m1.342-11.351A28.92 28.92 0 0 0 39.15 4.5c-1.794 1.083-8.193 3.104-8.193 3.104s-4.841-2.534-8.294-2.732a53.462 53.462 0 0 1-8.125 2.37S8.613 4.5 5.498 4.974c1.693 2.674 2.946 8.541 2.946 8.541l-2.302 7.777s1.997 6.127 2.166 8.632a40.133 40.133 0 0 0-2.065 8.769s.971-.046 2.394.009"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m19.906 27.434l3.114 3.115l-12.952 12.952l-3.114-3.115zm5.699-.858l-3.278 3.278l-1.727-1.726l3.238-3.238m6.923-3.792a.949.949 0 1 1 .332-1.3a.949.949 0 0 1-.332 1.3Zm4.764-2.825a.949.949 0 1 1 .332-1.3a.949.949 0 0 1-.332 1.3Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.149 14.651h0a6.783 6.783 0 0 1 9.294 2.374h0a.964.964 0 0 1-.338 1.322h0L29.07 24.299a.965.965 0 0 1-1.322-.338h0a6.783 6.783 0 0 1 2.374-9.294Zm-5.554 2.118l2.814.963m7.262-6.938l-.505 2.931"/>`),
		g.Group(children),
	)
}