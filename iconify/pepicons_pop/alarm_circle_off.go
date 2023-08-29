package pepicons_pop

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AlarmCircleOff(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 26 26"),
		g.Raw(`<g fill="currentColor"><path fill-rule="evenodd" d="M13 8.5a5.5 5.5 0 1 0 0 11a5.5 5.5 0 0 0 0-11ZM5.5 14a7.5 7.5 0 1 1 15 0a7.5 7.5 0 0 1-15 0Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M5.793 21.207a1 1 0 0 0 1.414 0l1.5-1.5a1 1 0 1 0-1.414-1.414l-1.5 1.5a1 1 0 0 0 0 1.414Zm11.5-2.914a1 1 0 0 1 1.414 0l1.5 1.5a1 1 0 0 1-1.414 1.414l-1.5-1.5a1 1 0 0 1 0-1.414ZM13 10a1 1 0 0 1 1 1v3a1 1 0 1 1-2 0v-3a1 1 0 0 1 1-1Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M17 14a1 1 0 0 1-1 1h-3a1 1 0 1 1 0-2h3a1 1 0 0 1 1 1ZM9.332 6.172L5.574 9.76s-.077-.07-.19-.198c-.257-.292-.697-.881-.821-1.601c-.113-.654.035-1.417.818-2.165c.783-.747 1.551-.86 2.2-.717c.713.157 1.281.624 1.561.894c.123.118.19.199.19.199Zm7.336 0l3.758 3.588s.077-.07.19-.198c.257-.292.697-.881.821-1.601c.113-.654-.035-1.417-.818-2.165c-.783-.747-1.551-.86-2.2-.717c-.713.157-1.281.624-1.561.894c-.123.118-.19.199-.19.199Z" clip-rule="evenodd"/><path d="M4.293 5.707a1 1 0 0 1 1.414-1.414l16 16a1 1 0 0 1-1.414 1.414l-16-16Z"/><path fill-rule="evenodd" d="M13 24c6.075 0 11-4.925 11-11S19.075 2 13 2S2 6.925 2 13s4.925 11 11 11Zm0 2c7.18 0 13-5.82 13-13S20.18 0 13 0S0 5.82 0 13s5.82 13 13 13Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}