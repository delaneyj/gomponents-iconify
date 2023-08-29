package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func OdontologyImplantNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsOdontologyImplantNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm22.796 30.753a134.222 134.222 0 0 1-5.42-.174l-2.78 1.001l-.03-.13a1.986 1.986 0 0 1 .042-1.06c-3.027-.256-5.17-.634-5.383-1.104c0-2.564-.855-5.232-1.685-7.82c-1.368-4.263-2.665-8.307.05-11.301c4.362-4.812 9.73-5.841 15.91-1.031l.03-.02l6.817 4.962a1 1 0 1 0 1.177-1.617l-6.086-4.43c3.16-1.57 9.53-3.751 14.14 1.105c4.12 4.341 2.14 8.877.233 13.242c-1.036 2.373-2.05 4.696-2.05 6.91c-.483.846-4.652 1.296-9.663 1.433l-12.65 4.554l-.4-1.73l7.748-2.79ZM16.3 38.966l16.795-6.046l.34-1.47a2.01 2.01 0 0 0 .05-.544L15.9 37.236l.399 1.73Zm16.26-3.728l-15.547 5.597a3.001 3.001 0 0 0 2.027 1.145l13.049-4.698l.471-2.044ZM28.613 42H24.89l6.665-2.4l-.017.075A3 3 0 0 1 28.613 42Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsOdontologyImplantNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}