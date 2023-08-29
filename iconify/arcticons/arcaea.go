package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Arcaea(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.856 6.063c-1.763 3.97-8.912 12.045-13.58 12.886c0 0 5.088 2.27 12.196-2.187a3.954 3.954 0 0 0-.379 2.523s6.106-4.316 7.196-8.067"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13 10.53c4.32 2.272 11.282 3.599 20.606-.842m1.704 8.668c3.073.298 7.53-1.889 7.53-1.889a4.655 4.655 0 0 0 2.193 3.06"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.113 4.418c-1.795 3.26-6.632 16.591 5.438 20.67c-6.939.337-12.136-6.73-9.664-18.127"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.275 21.947c-1.009 2.427-7.414 9.345-9.32 9.84m41.329-4.733a7.555 7.555 0 0 1-11.678 1.568c.774 3.784 4.234 4.878 8.44 3.448a24.694 24.694 0 0 0-10.733 12.154"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M32.398 41.79s-6.253 1.635-10.122 1.635c-1.683-2.607-7.052-11.565-7.584-14.145a17.939 17.939 0 0 1 .114-5.9m9.643 6.966a4.734 4.734 0 0 1 3.448 1.514m8.832-12.333c.449 1.903 3.365 5.375 5.776 5.573a2.776 2.776 0 0 0 2.975-2.017m-6.673-5.13c1.791.393 5.275 1.409 6.477 2.99m-26.746-8.575a9.424 9.424 0 0 0-3.131.819c.546 2.019 2.817 4.542 5.65 4.57s4.822-1.318 5.326-2.285a6.468 6.468 0 0 0-2.886-2.91M6.957 29.59s6.557 5.158 9.445 6.532c-.056 1.514.308 5.887.308 5.887a32.625 32.625 0 0 1 4.655 3.331m-9.658-20.96a24.306 24.306 0 0 0 2.774 3.632m15.141-7.069a2.025 2.025 0 0 0-1.528 2.211"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M40.885 18.51c-1.703.291-2.623 2.703-2.005 4.441M18.87 12.42a4.155 4.155 0 0 0 .949 5.164M4.738 22.65c.757-2.776 5.691-7.587 5.691-7.587M5.53 12.99a19.087 19.087 0 0 0 4.732-.98m33.44 7.67c.356.9.359 1.435-.398 1.792a1.107 1.107 0 0 1-1.64-.741c-.21-.983.339-1.837.658-1.705m-20.846-6.399c-.42 1.464.45 2.824 2.139 2.375s1.423-1.36 1.423-1.36"/>`),
		g.Group(children),
	)
}