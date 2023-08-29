package arcticons

import (
	g "github.com/maragudk/gomponents"
	s "github.com/maragudk/gomponents/svg"
)

func UobTmrw(children ...g.Node) g.Node {
	return s.SVG(
		g.Attr("viewbox", "0 0 48 48"),
		g.Raw(`<rect width="37" height="37" x="5.5" y="5.5" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" rx="2" ry="2"/><path fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" d="M29.271 23.608h0c-1.57 0-2.724-1.257-2.724-2.725v-2.828c0-1.572 1.258-2.829 2.724-2.829h0a2.817 2.817 0 0 1 2.83 2.804v2.749a2.817 2.817 0 0 1-2.806 2.829h-.024Zm8.343-4.191h-3.563m3.353 0c1.153 0 2.096.943 2.096 2.096a2.101 2.101 0 0 1-2.096 2.095h-3.457v-8.381h3.457c1.153 0 2.096.943 2.096 2.095s-.943 2.096-2.096 2.096v-.001Zm-18.151-4.191v5.553c0 1.572 1.257 2.829 2.724 2.829c1.466 0 2.724-1.257 2.724-2.829v-5.553M8.5 19.417h8.627m-1.507-4.191v8.382m-1.87-8.382v8.382m-1.872-8.382v8.382m-1.871-8.382v8.382m25.669 3.941l-1.307 5.225l-1.305-5.224l-1.306 5.224l-1.306-5.224m-16.43 5.224v-5.225m-1.697 0h3.395m11.371 3.559l1.601 1.6m-3.395.066V27.55h1.698c.98 0 1.763.784 1.763 1.764a1.755 1.755 0 0 1-1.748 1.763h-1.712m-7.497 1.631v-5.159l2.611 5.225l2.612-5.225v5.225"/>`),
		g.Group(children),
	)
}