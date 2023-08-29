package openmoji

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func CardIndex(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 72 72"),
		g.Raw(`<g stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="#ea5a47" d="M18.6 7h5.87a1.57 1.57 0 0 1 1.56 1.57V12h-9V8.57A1.57 1.57 0 0 1 18.6 7z"/><path fill="#b1cc33" d="M30 7h5.87a1.57 1.57 0 0 1 1.56 1.57V12h-9V8.57A1.57 1.57 0 0 1 30 7z"/><path fill="#92d3f5" d="M45 7h5.87a1.57 1.57 0 0 1 1.56 1.57V12h-9V8.57A1.57 1.57 0 0 1 45 7z"/><path fill="#d0cfce" d="M17 12h38v22H17zm0 26h38v22H17z"/></g><g stroke="#000" stroke-linecap="round" stroke-linejoin="round" stroke-width="2"><path fill="none" d="M18.6 7h5.87a1.57 1.57 0 0 1 1.56 1.57V12h0h-9h0V8.57A1.57 1.57 0 0 1 18.6 7zM30 7h5.87a1.57 1.57 0 0 1 1.56 1.57V12h0h-9h0V8.57A1.57 1.57 0 0 1 30 7zm15 0h5.87a1.57 1.57 0 0 1 1.56 1.57V12h0h-9h0V8.57A1.57 1.57 0 0 1 45 7zm-28 5h38v22H17zm0 26h38v22H17zm11-7v10m16-10v10"/><path d="M12 32h2v8h-2zm46 0h2v8h-2z"/><path fill="none" d="M22 18h10m-10 4h19m-19 4h10"/></g>`),
		g.Group(children),
	)
}