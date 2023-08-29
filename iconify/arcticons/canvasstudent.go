package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Canvasstudent(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 33.13a2.13 2.13 0 1 0 2.13 2.13A2.13 2.13 0 0 0 24 33.13Zm0 6.3a6.08 6.08 0 0 0-6 5.21a21.23 21.23 0 0 0 12 0a6.08 6.08 0 0 0-6-5.21Z"/><circle cx="16.04" cy="31.96" r="2.13" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.09 34.91a6 6 0 0 0-7.92-.54a21.64 21.64 0 0 0 8.46 8.46a6 6 0 0 0-.54-7.92Z"/><circle cx="12.74" cy="24" r="2.13" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M8.57 24a6.08 6.08 0 0 0-5.21-6a21.23 21.23 0 0 0 0 12a6.08 6.08 0 0 0 5.21-6Z"/><circle cx="16.04" cy="16.04" r="2.13" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.09 13.09a6 6 0 0 0 .54-7.92a21.64 21.64 0 0 0-8.46 8.46a6 6 0 0 0 7.92-.54ZM24 10.61a2.14 2.14 0 1 0 2.13 2.15v0A2.13 2.13 0 0 0 24 10.61Zm0-2.03a6.06 6.06 0 0 0 6-5.22a21.23 21.23 0 0 0-12 0a6.06 6.06 0 0 0 6 5.22Z"/><circle cx="31.95" cy="16.04" r="2.13" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.91 13.09a6 6 0 0 0 7.92.54a21.64 21.64 0 0 0-8.46-8.46a6 6 0 0 0 .54 7.92Z"/><circle cx="35.26" cy="24" r="2.13" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M44.64 18a6 6 0 0 0 0 12a21.23 21.23 0 0 0 0-12Z"/><circle cx="31.95" cy="31.96" r="2.13" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M34.91 34.91a6 6 0 0 0-.54 7.92a21.64 21.64 0 0 0 8.46-8.46a6 6 0 0 0-7.92.54Z"/>`),
		g.Group(children),
	)
}