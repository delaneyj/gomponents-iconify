package healthicons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func WaterTreatment(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<g fill="currentColor"><path d="M30 26.818C30 30.232 27.314 33 24 33s-6-2.768-6-6.182C18 21.41 24 15 24 15s6 6.41 6 11.818Z"/><path fill-rule="evenodd" d="M24 31c2.153 0 4-1.816 4-4.182c0-2.08-1.206-4.613-2.67-6.838A30.846 30.846 0 0 0 24 18.12a30.935 30.935 0 0 0-1.33 1.86C21.207 22.206 20 24.737 20 26.819C20 29.184 21.847 31 24 31Zm-1.303-14.46C20.932 18.767 18 23.037 18 26.818C18 30.232 20.686 33 24 33s6-2.768 6-6.182c0-3.78-2.932-8.05-4.697-10.278C24.543 15.58 24 15 24 15s-.544.58-1.303 1.54Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M14.86 10.867A16 16 0 0 0 28.784 39.27l.598 1.908A18 18 0 0 1 13.718 9.225l1.143 1.642Zm18.28 26.266A16 16 0 0 0 19.216 8.732l-.598-1.909a18 18 0 0 1 15.663 31.952l-1.143-1.642Z" clip-rule="evenodd"/><path fill-rule="evenodd" d="M14 10H7V8h9v9h-2v-7Zm20 28h7v2h-9v-9h2v7Z" clip-rule="evenodd"/></g>`),
		g.Group(children),
	)
}