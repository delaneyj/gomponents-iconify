package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Acc(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M31.253 11.902a16.384 16.384 0 0 0-6.966-2.57a16.659 16.659 0 1 0-4.004 33.075a16.857 16.857 0 0 0 5.295-.25a3.803 3.803 0 0 0 1.893-6.41l-1.17-1.17a2.583 2.583 0 0 1 .332-3.932a2.678 2.678 0 0 1 3.417.38l1.602 1.601a3.81 3.81 0 0 0 6.31-1.527a16.702 16.702 0 0 0 .675-7.669a16.376 16.376 0 0 0-2.684-6.907"/><circle cx="17.312" cy="17.305" r="2.804" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="15.253" cy="32.905" r="2.804" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="12.698" cy="24.635" r="2.804" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.774 23.416a4.214 4.214 0 0 0-8.017-1.822l-.01-.012s-1.776 3.318-3.666 3.584c.719 1.544 4.056 3.853 8.416 2.365l-.005-.006a4.219 4.219 0 0 0 3.282-4.11Z"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M25.512 19.304c.885-3.507 13.607-16.15 16.561-13.434s-10.29 15.559-13.475 16.336"/>`),
		g.Group(children),
	)
}