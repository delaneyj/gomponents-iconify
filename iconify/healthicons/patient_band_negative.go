package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func PatientBandNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsPatientBandNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM17 7a3 3 0 0 0-3 3v8a2 2 0 0 0-2 2H6v2h6v8H6v2h6a2 2 0 0 0 1.041 1.755c.484.266.959.693.959 1.245a1 1 0 0 0 1 1h4a1 1 0 0 0 1-1c0-.552.474-.98.959-1.245a2.01 2.01 0 0 0 .674-.6l4.229 2.33c3.947 2.174 9.276 1.888 12.806-1.019A16.04 16.04 0 0 0 40.222 33h.278a2.5 2.5 0 0 0 2.5-2.5v-1a2.49 2.49 0 0 0-.162-.888A2.498 2.498 0 0 0 44 26.5v-1c0-.889-.464-1.669-1.162-2.112A2.49 2.49 0 0 0 43 22.5v-1a2.5 2.5 0 0 0-2-2.45v-.81a3 3 0 0 0-2.767-2.992L35.039 15h-.653a31.143 31.143 0 0 0-13.638 3.145A1.995 1.995 0 0 0 20 18v-8a3 3 0 0 0-3-3Zm-1.5 24a.5.5 0 0 1-.5-.5v-9a.5.5 0 0 1 .5-.5h3a.5.5 0 0 1 .5.5v9a.5.5 0 0 1-.5.5h-3Zm11.327 2.733L22 31.073V20a2.02 2.02 0 0 0-.013-.23A29.143 29.143 0 0 1 34.386 17h.575l3.117.242a1 1 0 0 1 .922.997V19h-2v2h3.5a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5H37v2h4.5a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5H37v2h3.5a.5.5 0 0 1 .5.5v1a.5.5 0 0 1-.5.5h-8.447v2H37.3c-2.85 2.272-7.223 2.524-10.474.733ZM17 11.555a1 1 0 1 0 0-2a1 1 0 0 0 0 2ZM18 14a1 1 0 1 1-2 0a1 1 0 0 1 2 0Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsPatientBandNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}