package ps

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func Home(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 512 512"),
		g.Raw(`<path fill="currentColor" d="M256 7L6 257l30 30l28-28v226q0 18 12.5 30.5T107 528h298q18 0 30.5-12.5T448 485V259l30 30l30-30L271 25zm-43 478V336h86v149h-86zm128 0V336q0-17-12.5-30T299 293h-86q-17 0-29.5 13T171 336v149h-64V217L256 67l149 150v268h-64zm-85-341q-27 0-45.5 18.5T192 208t18.5 45.5T256 272t45.5-18.5T320 208t-18.5-45.5T256 144zm0 85q-21 0-21-21t21-21t21 21t-21 21z"/>`),
		g.Group(children),
	)
}