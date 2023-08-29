package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Gpscockpit(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M37.689 6.593A21.162 21.162 0 0 1 46.207 24a21.468 21.468 0 0 1-1.129 6.893m-1.301 3.125A20.967 20.967 0 0 1 24.707 45.5a21.436 21.436 0 0 1-8.668-1.819m-2.882-1.067A22.416 22.416 0 0 1 3.135 24.791m.329-3.514C4.574 10.446 13.582 2.5 24.707 2.5a21.422 21.422 0 0 1 9.978 2.45"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M19.409 9.08a13.638 13.638 0 0 1 5.756-.814a15.582 15.582 0 0 1 15.582 15.581a15.14 15.14 0 0 1-12.332 15.256m-3.25.326a15.262 15.262 0 0 1-11.594-5.108m-1.806-2.516a15.51 15.51 0 0 1-2.182-7.958a15.262 15.262 0 0 1 6.875-13.16"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M30.09 15.71a9.03 9.03 0 0 1 4.496 8.188a9.574 9.574 0 1 1-7.54-9.357"/><circle cx="17.792" cy="9.966" r="1.456" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="3.249" cy="23.085" r="1.456" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="14.436" cy="43.017" r="1.456" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="44.538" cy="32.542" r="1.456" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="36.3" cy="5.898" r="1.456" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="12.707" cy="33.051" r="1.456" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="28.47" cy="14.949" r="1.456" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="26.742" cy="39.091" r="1.456" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}