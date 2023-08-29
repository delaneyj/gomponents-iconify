package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func VespaMotorcycleNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsVespaMotorcycleNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM16.95 33a2.5 2.5 0 0 1-4.9 0h4.9Zm-7.928 0H4.47c-.26 0-.47-.18-.47-.4C4 27.85 8.516 24 14.086 24h8.07c.466 0 .844.322.844.72V26h6.875l.93-7.437c.502.278 1.08.437 1.695.437H35v8.759A5.499 5.499 0 0 1 37.933 27H38a5.5 5.5 0 1 1-5.475 6.032a.802.802 0 0 1-.02-.021l-.01-.011H21v-.002a2.556 2.556 0 0 1-.111.002h-.911a5.5 5.5 0 0 1-10.956 0ZM38 35a2.5 2.5 0 0 0 2.456-2.971l-4.925.868A2.5 2.5 0 0 0 38 35ZM9 21a3 3 0 0 1 3-3h9a3 3 0 0 1 3 3v1H9v-1Zm23.5-11c-1.286 0-2.41.694-3.018 1.727L25.346 11L25 12.97l4.004.704A3.5 3.5 0 0 0 32.5 17h3.885c.34 0 .615-.276.615-.615v-5.77a.615.615 0 0 0-.615-.615H32.5Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsVespaMotorcycleNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}