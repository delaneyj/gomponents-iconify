package emojione_monotone

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func JapaneseOpenForBusinessButton(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 64 64"),
		g.Raw(`<path fill="currentColor" d="M26.189 29.182h11.207v2.684H26.189zm-3.056 12.267H40.83v3.336H23.133z"/><path fill="currentColor" d="M52 2H12C6.477 2 2 6.477 2 12v40c0 5.523 4.477 10 10 10h40c5.523 0 10-4.477 10-10V12c0-5.523-4.477-10-10-10m-6.717 48H40.83v-1.266H23.133V50h-4.264V37.539h9.472c.188-.766.415-1.57.564-2.338h-6.867v-9.393h19.698v9.393h-7.963a27.58 27.58 0 0 1-.98 2.338h12.49V50M48 28.225h-4.34V23.93H20.076v4.295H16v-8.168h4.566c-.68-1.225-1.623-2.607-2.49-3.756l3.887-1.688c1.244 1.457 2.717 3.49 3.358 4.908l-1.208.535h5.813c-.529-1.379-1.548-3.258-2.454-4.715L31.396 14c1.095 1.496 2.341 3.682 2.867 5.061l-2.641.996h6.188c1.208-1.646 2.642-3.986 3.358-5.787L46 15.803c-1.095 1.457-2.227 2.99-3.283 4.254H48v8.168"/>`),
		g.Group(children),
	)
}