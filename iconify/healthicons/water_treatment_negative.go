package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterTreatmentNegative(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="none"><g clip-path="url(#healthiconsWaterTreatmentNegative0)"><path fill="currentColor" fill-rule="evenodd" d="M48 0H0v48h48V0ZM24 33c3.314 0 6-2.768 6-6.182C30 21.41 24 15 24 15s-6 6.41-6 11.818C18 30.232 20.686 33 24 33ZM7 10h5.686a18 18 0 0 0 16.695 31.177l-.598-1.909A16 16 0 0 1 14 11.51V17h2V8H7v2Zm28.314 28H41v2h-9v-9h2v5.49A16 16 0 0 0 19.217 8.732l-.598-1.909A18 18 0 0 1 35.314 38Z" clip-rule="evenodd"/></g><defs><clipPath id="healthiconsWaterTreatmentNegative0"><path d="M0 0h48v48H0z"/></clipPath></defs></g>`),
		g.Group(children),
	)
}