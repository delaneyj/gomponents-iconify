package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Galaxyinvaders(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M24 19.364V5.5m-1.3 13.43l2.6-2.6l-2.6-2.6l2.6-2.599l-2.6-2.6l2.6-2.6m-.434 15.166a16.882 16.882 0 0 1 1.733 3.717a9.105 9.105 0 0 0 1.746 3.75c1.313 2.055 9.952-9.633 9.952-9.633l-4.765 9.531l7.798-3.466l-8.232 6.932l9.099.434l-10.398 1.3l6.065 6.498l-7.798-4.766l2.6 7.106l-4.595-5.754s-.82-3.97-1.094-4.67a1.23 1.23 0 0 0-.584-.592c-.05.41-.086.818-.145 1.236l-4.495-.075c-.051-.414-.117-.822-.145-1.241a1.174 1.174 0 0 0-.584.575c-.288.687-1.094 4.65-1.094 4.65L15.335 42.5l2.6-7.105l-7.8 4.765l6.066-6.499l-10.398-1.3l9.099-.433l-8.232-6.931l7.799 3.466l-4.766-9.532s8.593 11.75 9.952 9.568a9.56 9.56 0 0 0 1.745-3.743a17.729 17.729 0 0 1 1.734-3.659Zm-.29 15.014A5.643 5.643 0 0 0 24 38.713a6.315 6.315 0 0 0-.577-2.638c-.302-.391-.844-.695-.77-1.324a1.35 1.35 0 0 1 2.694 0c.04.649-.426.96-.77 1.36Z"/><circle cx="24" cy="27.163" r="1.3" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.2 15.898v-2.166m0 12.564V24.13m0-3.033v-2.166m-10.4-3.033v-2.166m0 12.564V24.13m0-3.033v-2.166m17.764-5.372l.434-3.033m-1.3 9.098l.433-3.033m-24.695-3.032l-.434-3.033m1.3 9.098l-.433-3.033"/>`),
		g.Group(children),
	)
}