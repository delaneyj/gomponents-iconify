package solar

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func LightningBoldDuotone(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 24 24"),
		g.Raw(`<defs><path id="solarLightningBoldDuotone0" fill-rule="evenodd" d="M10.8 8h2.4c2.263 0 3.394 0 4.097.692C18 9.385 18 10.5 18 12.727v2.546c0 2.228 0 3.342-.703 4.035C16.594 20 15.463 20 13.2 20h-2.4c-2.263 0-3.394 0-4.097-.692C6 18.615 6 17.5 6 15.273v-2.546c0-2.228 0-3.342.703-4.035C7.406 8 8.537 8 10.8 8Zm1.8 4.45a.75.75 0 1 0-1.2-.9l-1.5 2a.75.75 0 0 0 .6 1.2H12l-.6.8a.75.75 0 1 0 1.2.9l1.5-2a.75.75 0 0 0-.6-1.2H12l.6-.8Z" clip-rule="evenodd"/></defs><g fill="currentColor"><path fill-rule="evenodd" d="M12.5 19.25a.75.75 0 0 1 .75.75v2a.75.75 0 0 1-1.5 0v-2a.75.75 0 0 1 .75-.75Z" clip-rule="evenodd" opacity=".5"/><use href="#solarLightningBoldDuotone0" fill-rule="evenodd" clip-rule="evenodd"/><path d="M10.25 4.75A.75.75 0 0 1 11 4h2a.75.75 0 0 1 0 1.5h-2a.75.75 0 0 1-.75-.75Z"/><use href="#solarLightningBoldDuotone0" fill-rule="evenodd" clip-rule="evenodd"/><path d="M8.5 5.5c0-1.404 0-2.107.337-2.611a2 2 0 0 1 .552-.552C9.893 2 10.596 2 12 2c1.405 0 2.107 0 2.611.337a2 2 0 0 1 .552.552c.337.504.337 1.207.337 2.611v2.55C14.898 8 14.148 8 13.2 8h-2.4c-.948 0-1.697 0-2.3.05V5.5Z" opacity=".5"/></g>`),
		g.Group(children),
	)
}