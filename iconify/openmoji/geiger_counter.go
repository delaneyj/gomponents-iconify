package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func GeigerCounter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<rect width="36" height="49" x="16" y="12" fill="#f4aa41" rx="6"/><circle cx="34" cy="26" r="10" fill="#fff"/><ellipse cx="34" cy="31.5" fill="#d0cfce" rx="8.348" ry="4.5"/><path fill="#e27022" d="M52 55V18a6.018 6.018 0 0 0-6-6v49a6.018 6.018 0 0 0 6-6Z"/><path fill="#f1b31c" d="M32 31h4v25h-4z"/><g stroke="#000" stroke-linecap="round" stroke-linejoin="round"><path d="M24.754 45.133a.629.629 0 1 1-.315-.545a.629.629 0 0 1 .315.545Zm-3.146 0a2.515 2.515 0 0 1 1.258-2.18l-1.258-2.18a5.03 5.03 0 0 0-2.516 4.36Zm3.775 2.18a2.52 2.52 0 0 1-2.517 0l-1.258 2.18a5.04 5.04 0 0 0 5.033 0Zm0-4.36a2.515 2.515 0 0 1 1.258 2.18h2.516a5.03 5.03 0 0 0-2.516-4.359Z"/><rect width="36" height="49" x="16" y="12" fill="none" stroke-width="2" rx="6"/><path fill="none" stroke-width="2" d="M42.347 31.5a9.992 9.992 0 0 0-16.695 0M29 23l1.797 4.494"/><path fill="none" stroke-width="2" d="M36 35.8a10 10 0 1 0-4 0m0 14.755a4 4 0 1 0 4 0"/><rect width="4" height="25" x="32" y="31" fill="none" stroke-width="2" rx="1"/></g>`),
		g.Group(children),
	)
}