package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PetBreathCounter(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<circle cx="28.242" cy="16.718" r="3.786" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M35.455 23.778c1.382-1.126 3.99-3.357 2.916 2.26c5.551 2.306 3.313 4.275 3.821 5.225a1.974 1.974 0 0 0 1.206.982M38.79 29.28c.716.073 1.046.162 1.513.566m-1.844 2.786c.382 1.697 2.426 3.871 5.041.975m-11.23-7.425c-4.22 1.612-3.58 7.738-7.584 9.35"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M9.53 20.37c3.223-1.44 5.036-2.084 7.118-.384c2.609 2.13 4.931 7.942 13.036 8.377"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M15.938 23.17c.882 1.573 2.245 4.394 1.247 6.11m-2.225 3.233c.383 1.804.872 3.204 1.544 4.288m-5.184-8.074c.167 1.65.986 2.226 2.329 2.213m-6.625-8.288c-2.162.205-2.66.857-2.494 2.635s2.187 3.645 6.382.998m-.947-4.195c.34-.234.934-.192 1.266 0m17.011-5.372v-2.104m-1.202-3.415h2.404m-1.202 1.733v-1.733m2.733 2.901l.464-.465m-8.748 2.031h.345m-.345 2.138h.345"/>`),
		g.Group(children),
	)
}