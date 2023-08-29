package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func RunningWaterNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsRunningWaterNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm7 42V23H6V12h1V6h2v6h1v11H9v19H7ZM19 8h5a1 1 0 1 0 0-2H12a1 1 0 1 0 0 2h5v3h-5v13h13V11h-6V8Zm9 16v6h14v-6h-2v-3c0-5.523-4.477-10-10-10h-3v10h1a2 2 0 0 1 2 2v1h-2Zm7.5 18c1.933 0 3.5-1.628 3.5-3.636C39 35.182 35.5 32 35.5 32S32 35.182 32 38.364C32 40.372 33.567 42 35.5 42Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsRunningWaterNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}