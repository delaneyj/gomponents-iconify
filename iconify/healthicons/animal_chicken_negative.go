package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func AnimalChickenNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsAnimalChickenNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M0 0h48v48H0V0Zm19 42h5v-2.001a15.693 15.693 0 0 1-12.905-7.004c-1.885-2.84-1.632-6.46-.75-9.753l.76-2.842a8.225 8.225 0 0 0 .277-2.4H7.915C6.132 18 5.24 15.846 6.5 14.586L15.086 6c1.26-1.26 3.414-.368 3.414 1.414V9h1.37c1.839 0 2.703 2.273 1.328 3.495l-2.02 1.795l3.32 6.067a5 5 0 0 0 5.2 2.533l13.115-2.163c.788-.13 1.33.77.848 1.406c-1.164 1.536-1.71 3.428-2.078 5.32a15.27 15.27 0 0 1-1.113 3.455a.998.998 0 0 0-.22.187C36.148 33.479 33.003 35 29.478 35C23.102 35 18 30.039 18 24a1 1 0 1 0-2 0c0 7.216 6.072 13 13.478 13c1.863 0 3.639-.365 5.255-1.025A15.684 15.684 0 0 1 26 39.899V42h2.5v2H19v-2Zm-.807-29.51L19.87 11H18.5a2 2 0 0 1-2-2V7.414l-2.644 2.644a3.88 3.88 0 0 1 4.078 1.96l.26.472Zm-7.485.716L7.914 16h3.159l-.298-1.148a3.863 3.863 0 0 1-.067-1.646ZM14 14a1 1 0 1 0 0-2a1 1 0 0 0 0 2Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsAnimalChickenNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}