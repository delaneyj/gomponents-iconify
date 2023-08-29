package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterTreatmentOutline(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M41.425 28.515A18 18 0 0 1 35.314 38H41v2h-9v-9h2v5.49A16 16 0 0 0 19.217 8.732l-.598-1.909a18 18 0 0 1 22.806 21.692ZM8.511 19.987A16 16 0 0 1 14 11.51V17h2V8H7v2h5.686a18 18 0 0 0 16.695 31.177l-.598-1.908A16 16 0 0 1 8.51 19.987Z"/><path fill-rule="evenodd" d="M22.697 16.54C20.932 18.767 18 23.037 18 26.818C18 30.232 20.686 33 24 33s6-2.768 6-6.182c0-3.78-2.932-8.05-4.697-10.278C24.543 15.58 24 15 24 15s-.544.58-1.303 1.54ZM24 31c2.153 0 4-1.816 4-4.182c0-2.08-1.206-4.613-2.67-6.838A30.846 30.846 0 0 0 24 18.12a30.935 30.935 0 0 0-1.33 1.86C21.207 22.206 20 24.737 20 26.819C20 29.184 21.847 31 24 31Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}