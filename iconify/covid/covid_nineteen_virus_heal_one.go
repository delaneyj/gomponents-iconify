package covid

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CovidNineteenVirusHealOne(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<g fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="1.5"><path d="M23.25 17.25a1 1 0 0 0-1-1h-2.5v-2.5a1 1 0 0 0-1-1h-1.5a1 1 0 0 0-1 1v2.5h-2.5a1 1 0 0 0-1 1v1.5a1 1 0 0 0 1 1h2.5v2.5a1 1 0 0 0 1 1h1.5a1 1 0 0 0 1-1v-2.5h2.5a1 1 0 0 0 1-1v-1.5ZM4.183 12.754H2.162a1.413 1.413 0 0 1 0-2.825h2.021m5.746-5.746V2.162a1.413 1.413 0 0 1 2.825 0v2.021"/><path d="M15.405 5.28a7.3 7.3 0 0 0-8.128 0M18.5 9.929a7.25 7.25 0 0 0-1.1-2.652M7.277 17.4a7.271 7.271 0 0 0 2.478 1.061M5.28 7.277a7.3 7.3 0 0 0 0 8.128"/><path d="M6.848 9.511c0-.426-.17-.835-.471-1.137L4.518 6.515a1.414 1.414 0 0 1 2-2l1.856 1.862c.302.301.71.47 1.137.471m0 8.987a1.606 1.606 0 0 0-1.137.471l-1.859 1.859a1.414 1.414 0 0 1-2-2l1.859-1.86c.301-.301.47-.71.471-1.136m6.327-6.321c.426 0 .835-.17 1.136-.471l1.86-1.859a1.414 1.414 0 1 1 2 2l-1.862 1.856a1.606 1.606 0 0 0-.471 1.137"/></g>`),
		g.Group(children),
	)
}