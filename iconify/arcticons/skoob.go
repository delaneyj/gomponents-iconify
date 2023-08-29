package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Skoob(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M23.998 13.737s0-2.053-4.107-4.105L9.625 4.5v4.105a5.331 5.331 0 0 0 2.054 4.105M23.998 43.5a4.449 4.449 0 0 0-4.107-2.053c-3.08 0-6.16-1.026-7.187-4.105a5.985 5.985 0 0 1-.244-1.297"/><circle cx="16.274" cy="21.057" r="2.19" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24.002 21.256a1.368 1.368 0 0 0-1.367 1.369v0a1.368 1.368 0 0 0 .07.427l-.004-.002l.02.044a1.375 1.375 0 0 0 .072.162l1.21 2.738l1.208-2.738a1.375 1.375 0 0 0 .073-.162l.02-.044l-.004.002a1.359 1.359 0 0 0-1.296-1.796h-.002Zm-.004-7.519s0-2.053 4.107-4.105L38.373 4.5v4.105a5.331 5.331 0 0 1-2.054 4.105M23.998 43.5a4.45 4.45 0 0 1 4.107-2.053c3.08 0 6.16-1.026 7.188-4.105a5.992 5.992 0 0 0 .244-1.297"/><circle cx="31.723" cy="21.057" r="2.19" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M22.682 23.08a7.502 7.502 0 1 1 .895-1.766"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.407 23.211a7.505 7.505 0 1 0-1.12-2.316q.047.165.101.326"/><path fill="none" stroke="currentColor" stroke-linejoin="round" d="m25.732 32.964l15.338-8.852c.2 4.26-1.735 7.416-5.547 9.604L24 40.366l-11.523-6.65c-3.811-2.188-5.746-5.345-5.547-9.604l15.338 8.852A3.529 3.529 0 0 1 24 36.204a3.529 3.529 0 0 1 1.732-3.24Z"/>`),
		g.Group(children),
	)
}