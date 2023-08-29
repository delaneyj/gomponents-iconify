package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Socratic(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M13.088 22.496V4.5l3.737 3.738H33.3l3.485-3.486v6.992m0 7.75v7.345a11.779 11.779 0 0 1-11.802 11.63H13.088V35.32"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="m13.088 22.334l-6.11 13.075s4.927 2.856 8.197-2.795l2.005.706s2.922-7.83-4.093-10.988Zm23.698 4.505l4.237 8.695s-4.13 1.943-7.018-1.232M18.816 38.47v5.03H21.6m3.383-5.03v5.03h2.783m1.842-22.169l-3.553 2.052V19.28Z"/><circle cx="20.601" cy="15.596" r="5.006" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="33.283" cy="15.596" r="5.006" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M28.279 15.596h-2.693m-9.99 0h-2.508a2.187 2.187 0 0 0-2.23 2.231m15.03 9.141a1.761 1.761 0 1 1-3.523 0m8.991 0a1.761 1.761 0 1 1-3.522 0m.954 3.553a1.761 1.761 0 1 1-3.523 0"/><circle cx="20.601" cy="15.596" r="1.873" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><circle cx="33.283" cy="15.596" r="1.873" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/>`),
		g.Group(children),
	)
}