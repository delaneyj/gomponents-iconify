package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Supertuxkart(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="24" cy="24" r="21.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.947 19.96a4.536 4.536 0 0 1 2.1 2.092l10.702-2.81a15.528 15.528 0 0 0-9.99-9.99Zm-5.987 2.093a4.537 4.537 0 0 1 2.092-2.1L19.241 9.252a15.528 15.528 0 0 0-9.99 9.989Zm8.08 3.894a4.536 4.536 0 0 1-2.092 2.1l2.811 10.702a15.528 15.528 0 0 0 9.99-9.99Zm-5.988 2.1a4.537 4.537 0 0 1-2.092-2.1L9.251 28.759a15.528 15.528 0 0 0 9.989 9.99Z"/><circle cx="24" cy="24" r="4.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}